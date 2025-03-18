package booksrepo

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"github.com/srijilv/go-api-template.git/db"
	"github.com/srijilv/go-api-template.git/pkg/components"
	"github.com/srijilv/go-api-template.git/pkg/domain/model/books"
	"github.com/srijilv/go-api-template.git/pkg/domain/model/books/data"
	apierror "github.com/srijilv/go-api-template.git/utils/api_error"
	"go.uber.org/multierr"
)

const (
	PgsqlDBFindConnErr apierror.Code = iota + 290
	PgsqlDBFindInitErr
	PgsqlDBFindFinErr

	ErrListBooksNotFound apierror.Code = iota + 300
	ErrListBooksUnknown
	ErrListBooksScan
)

var ListBooksQuery string = fmt.Sprintf(`SELECT id, 
											title, 
											author, 
											genre, 
											published_year, 
											isbn, 
											price, 
											status, 
											created_by, 
											created_at, 
											updated_by, 
											updated_at 
											FROM %s 
											WHERE status = $1 
											ORDER BY id DESC
											LIMIT $2 OFFSET $3;`, booksTable)

func (p PgsqlBookRepository) ListBooks(ctx context.Context, page, limit int32) (records []books.Books, err error) {

	offset := (page - 1) * limit
	pool, err := PgSqlPoolConnection(ctx, p.poolProvider, p.dbName, PgsqlDBFindConnErr)
	if err != nil {
		return
	}

	defer func() {
		PgSqlPoolClose(pool)
	}()

	tx, err := PgSqlTransactionBegin(ctx, pool, PgsqlDBFindInitErr)
	if err != nil {
		return
	}

	defer func() {
		err = PgSqlTransactionReadOnly(ctx, tx, err, PgsqlDBFindFinErr)
		if err != nil {
			return
		}
	}()

	rows, err := tx.Query(ctx, ListBooksQuery, "active", limit, offset)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
			"errorCode": ErrListBooksNotFound,
			"error":     err,
		}).Errorf("error occurred while fetching data")
		err = apierror.NewUnknownError(ErrListBooksNotFound, components.ApiComponent, err)
		return
	}

	fmt.Printf("rows: %+v\n", rows.RawValues())

	defer rows.Close()

	var (
		loadable  data.LoadableBooks
		loadables []data.LoadableBooks

		updatedBy pgtype.Int8
	)

	for rows.Next() {
		if err = rows.Scan(&loadable.Id, &loadable.Title, &loadable.Author, &loadable.Genre,
			&loadable.PublishedYear, &loadable.Isbn, &loadable.Price, &loadable.Status, &loadable.CreatedBy,
			&loadable.CreatedAt, &updatedBy, &loadable.UpdatedAt); err != nil {
			logrus.WithFields(logrus.Fields{
				"component": components.ApiComponent,
				"layer":     layer,
				"errorCode": ErrListBooksScan,
				"error":     err,
			}).Errorf("error occurred while scaning data")
			err = apierror.NewUnknownError(ErrListBooksScan, components.ApiComponent, err)
			return
		}
		loadable.UpdatedBy = updatedBy.Int
		loadables = append(loadables, loadable)
	}

	if len(loadables) == 0 {
		logrus.WithFields(logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
			"errorCode": ErrListBooksNotFound,
			"error":     err,
		}).Errorf("no data found")
		err = apierror.NewNotFoundError(ErrListBooksNotFound, components.ApiComponent, fmt.Errorf("data not found"))
		return
	}

	// switch err {
	// case pgx.ErrNoRows:
	// 	logrus.WithFields(logrus.Fields{
	// 		"component": components.ApiComponent,
	// 		"layer":     layer,
	// 		"errorCode": ErrListBooksNotFound,
	// 		"error":     err,
	// 	}).Errorf("not found")
	// 	err = apierror.NewNotFoundError(ErrListBooksNotFound, components.ApiComponent, err)
	// default:
	// 	logrus.WithFields(logrus.Fields{
	// 		"component": components.ApiComponent,
	// 		"layer":     layer,
	// 		"errorCode": ErrListBooksNotFound,
	// 		"error":     err,
	// 	}).Errorf("error occurred while fetching data")
	// 	err = apierror.NewUnknownError(ErrListBooksNotFound, components.ApiComponent, err)
	// }

	return books.UnmarshalSlice(loadables)
}

func PgSqlPoolConnection(ctx context.Context, poolProvider db.PgSqlConnectionPoolProvider, db string, errCode apierror.Code) (pool db.PgxIface, err error) {
	pool, err = poolProvider.GetPgSqlConnectionPool(ctx, db)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
			"dbType":    dbType,
			"errorCode": errCode,
			"error":     err,
		}).Error("pool cannot be connection error")
		err = apierror.NewUnknownError(errCode, components.ApiComponent, err)
	}

	return
}

func PgSqlPoolClose(pool db.PgxIface) {
	pool.Close()
}

func PgSqlTransactionBegin(ctx context.Context, p db.PgxIface, errCode apierror.Code) (tx pgx.Tx, err error) {
	tx, err = p.Begin(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
			"dbType":    dbType,
			"errorCode": errCode,
			"error":     err,
		}).Error("database transaction cannot be started")
		err = apierror.NewUnknownError(errCode, components.ApiComponent, err)
	}

	return
}

func PgSqlTransactionReadOnly(ctx context.Context, tx pgx.Tx, e error, errCode apierror.Code) (err error) {
	err = FinishTransactionReadOnly(ctx, tx, e)
	if err != nil {
		if _, ok := err.(apierror.ApiError); ok {
			return
		}
		logrus.WithFields(logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
			"dbType":    dbType,
			"errorCode": errCode,
			"error":     err,
		}).Error("database transaction cannot be finished")
		err = apierror.NewUnknownError(errCode, components.ApiComponent, err)
	}

	return
}

func PgSqlTransaction(ctx context.Context, tx pgx.Tx, e error, errCode apierror.Code) (err error) {
	err = FinishTransaction(ctx, tx, e)
	if err != nil {
		if _, ok := err.(apierror.ApiError); ok {
			return
		}
		logrus.WithFields(logrus.Fields{
			"component": components.ApiComponent,
			"layer":     layer,
			"dbType":    dbType,
			"errorCode": errCode,
			"error":     err,
		}).Error("database transaction cannot be finished")
		err = apierror.NewUnknownError(errCode, components.ApiComponent, err)
	}

	return
}

func FinishTransaction(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		logrus.Errorf("rolling back transaction because of error during transactions: %v", err)
		if rollBackErr := tx.Rollback(ctx); rollBackErr != nil {
			logrus.Errorf("rolling back failed: %v", err)
			rollBackErr = fmt.Errorf("rolling back failed: %w", rollBackErr)
			err = multierr.Append(err, rollBackErr)
		}
		return err
	}

	logrus.Infof("committing transaction...")
	if commitErr := tx.Commit(ctx); commitErr != nil {
		logrus.Errorf("committing failed: %v", err)
		err = fmt.Errorf("committing failed: %w", commitErr)
	}
	return err
}

func FinishTransactionReadOnly(ctx context.Context, tx pgx.Tx, err error) error {
	rollBackErr := tx.Rollback(ctx)
	if rollBackErr != nil {
		rollBackErr = fmt.Errorf("rolling back failed: %w", rollBackErr)
		if err != nil {
			err = multierr.Append(err, rollBackErr)
			return err
		}
		return rollBackErr
	}
	return err
}
