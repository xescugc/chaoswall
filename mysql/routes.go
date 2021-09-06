package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/cycloidio/sqlr"
	"github.com/xescugc/chaoswall/route"
	"golang.org/x/xerrors"
)

type RouteRepository struct {
	querier sqlr.Querier
}

type dbRoute struct {
	ID        sql.NullInt64
	Name      sql.NullString
	Canonical sql.NullString
}

func newDbRoute(w route.Route) *dbRoute {
	return &dbRoute{
		ID:        ToNullInt64(w.ID),
		Name:      ToNullString(w.Name),
		Canonical: ToNullString(w.Canonical),
	}
}

func (w *dbRoute) toDomainEntity() *route.Route {
	return &route.Route{
		ID:        ToUint32(w.ID),
		Name:      w.Name.String,
		Canonical: w.Canonical.String,
	}
}

func (w *dbRoute) scanFields() []interface{} {
	return []interface{}{
		&w.ID,
		&w.Name,
		&w.Canonical,
	}
}

func dbRouteFields(prefix string) string {
	fields := []string{"id", "name", "canonical"}
	s := ""
	for i, f := range fields {
		format := " %s.%s,"
		if i == len(fields)-1 {
			format = " %s.%s"
		}

		s += fmt.Sprintf(format, prefix, f)
	}
	return s
}

// NewRouteRepository returns an implementation of the route.Repository
func NewRouteRepository(db sqlr.Querier) route.Repository {
	return &RouteRepository{
		querier: db,
	}
}

func (r *RouteRepository) Create(ctx context.Context, gCan, wCan string, rr route.Route) (uint32, error) {
	dbr := newDbRoute(rr)
	res, err := r.querier.ExecContext(ctx, `
		INSERT INTO routes (name, canonical, wall_id)
		SELECT ?, ?, w.id
		FROM
			(
				SELECT w.id
				FROM walls AS w
				JOIN gyms AS g
					ON g.id = w.gym_id
				WHERE g.canonical = ? AND w.canonical = ?
			) AS w
	`, dbr.Name, dbr.Canonical, gCan, wCan)

	if err != nil {
		return 0, xerrors.Errorf("failed to ExecContext: %w", err)
	}

	id, err := LastInsertedID(res)
	if err != nil {
		return 0, xerrors.Errorf("failed to get LastInsertedID: %w", err)
	}

	return id, nil
}

func (r *RouteRepository) Filter(ctx context.Context, gCan, wCan string) ([]*route.Route, error) {
	rows, err := r.querier.QueryContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM routes AS r
		JOIN walls AS w
			ON w.id = r.wall_id
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ? AND w.canonical = ?
		ORDER BY r.id
		`, dbRouteFields("r")), gCan, wCan,
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to QueryContext: %w", err)
	}
	defer rows.Close()

	routes, err := scanRoutes(rows)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanRoutes: %w", err)
	}
	return routes, nil
}

func (r *RouteRepository) FindByCanonical(ctx context.Context, gCan, wCan, rCan string) (*route.Route, error) {
	row := r.querier.QueryRowContext(ctx, fmt.Sprintf(`
		SELECT %s
		FROM routes AS r
		JOIN walls AS w
			ON w.id = r.wall_id
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ? AND w.canonical = ? AND r.canonical = ?
	`, dbRouteFields("r")), gCan, wCan, rCan)

	w, err := scanRoute(row)
	if err != nil {
		return nil, xerrors.Errorf("failed to scanRoute: %w", err)
	}

	return w, nil
}

func (r *RouteRepository) UpdateByCanonical(ctx context.Context, gCan, wCan, rCan string, rr route.Route) error {
	dbr := newDbRoute(rr)
	_, err := r.querier.ExecContext(ctx, `
		UPDATE routes AS r
		JOIN walls AS w
			ON w.id = r.wall_id
		JOIN gyms AS g
			ON g.id = w.gym_id
		SET r.name = ?, r.canonical = ?
		WHERE g.canonical = ? AND w.canonical = ? AND r.canonical = ?
	`, dbr.Name, dbr.Canonical, gCan, wCan, rCan)

	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

func (r *RouteRepository) DeleteByCanonical(ctx context.Context, gCan, wCan, rCan string) error {
	_, err := r.querier.ExecContext(ctx, `
		DELETE r
		FROM routes AS r
		JOIN walls AS w
			ON w.id = r.wall_id
		JOIN gyms AS g
			ON g.id = w.gym_id
		WHERE g.canonical = ? AND w.canonical = ? AND r.canonical = ?
	`, gCan, wCan, rCan)
	if err != nil {
		return xerrors.Errorf("failed to ExecContext: %w", err)
	}

	return nil
}

// scanRoute scans and returns a Route from a row
func scanRoute(s sqlr.Scanner) (*route.Route, error) {
	var g dbRoute

	err := s.Scan(g.scanFields()...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("route not found")
		}
		return nil, xerrors.Errorf("failed to Scan: %w", err)
	}

	return g.toDomainEntity(), nil
}

// scanRoutes scans and returns Routes from the rows,
// the caller is the one in charge of closing the rows
func scanRoutes(rows *sql.Rows) ([]*route.Route, error) {
	var gs []*route.Route

	for rows.Next() {
		g, err := scanRoute(rows)
		if err != nil {
			return nil, xerrors.Errorf("failed to scanRoute: %w", err)
		}
		gs = append(gs, g)
	}

	if err := rows.Err(); err != nil {
		return nil, xerrors.Errorf("failed to scan rows: %w", err)
	}

	return gs, nil
}
