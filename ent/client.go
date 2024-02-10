// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/atamayoz/credit-go/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/atamayoz/credit-go/ent/simulation"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Simulation is the client for interacting with the Simulation builders.
	Simulation *SimulationClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Simulation = NewSimulationClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Simulation: NewSimulationClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Simulation: NewSimulationClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Simulation.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Simulation.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Simulation.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *SimulationMutation:
		return c.Simulation.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// SimulationClient is a client for the Simulation schema.
type SimulationClient struct {
	config
}

// NewSimulationClient returns a client for the Simulation from the given config.
func NewSimulationClient(c config) *SimulationClient {
	return &SimulationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `simulation.Hooks(f(g(h())))`.
func (c *SimulationClient) Use(hooks ...Hook) {
	c.hooks.Simulation = append(c.hooks.Simulation, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `simulation.Intercept(f(g(h())))`.
func (c *SimulationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Simulation = append(c.inters.Simulation, interceptors...)
}

// Create returns a builder for creating a Simulation entity.
func (c *SimulationClient) Create() *SimulationCreate {
	mutation := newSimulationMutation(c.config, OpCreate)
	return &SimulationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Simulation entities.
func (c *SimulationClient) CreateBulk(builders ...*SimulationCreate) *SimulationCreateBulk {
	return &SimulationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SimulationClient) MapCreateBulk(slice any, setFunc func(*SimulationCreate, int)) *SimulationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SimulationCreateBulk{err: fmt.Errorf("calling to SimulationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SimulationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SimulationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Simulation.
func (c *SimulationClient) Update() *SimulationUpdate {
	mutation := newSimulationMutation(c.config, OpUpdate)
	return &SimulationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SimulationClient) UpdateOne(s *Simulation) *SimulationUpdateOne {
	mutation := newSimulationMutation(c.config, OpUpdateOne, withSimulation(s))
	return &SimulationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SimulationClient) UpdateOneID(id int) *SimulationUpdateOne {
	mutation := newSimulationMutation(c.config, OpUpdateOne, withSimulationID(id))
	return &SimulationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Simulation.
func (c *SimulationClient) Delete() *SimulationDelete {
	mutation := newSimulationMutation(c.config, OpDelete)
	return &SimulationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SimulationClient) DeleteOne(s *Simulation) *SimulationDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SimulationClient) DeleteOneID(id int) *SimulationDeleteOne {
	builder := c.Delete().Where(simulation.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SimulationDeleteOne{builder}
}

// Query returns a query builder for Simulation.
func (c *SimulationClient) Query() *SimulationQuery {
	return &SimulationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSimulation},
		inters: c.Interceptors(),
	}
}

// Get returns a Simulation entity by its id.
func (c *SimulationClient) Get(ctx context.Context, id int) (*Simulation, error) {
	return c.Query().Where(simulation.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SimulationClient) GetX(ctx context.Context, id int) *Simulation {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SimulationClient) Hooks() []Hook {
	return c.hooks.Simulation
}

// Interceptors returns the client interceptors.
func (c *SimulationClient) Interceptors() []Interceptor {
	return c.inters.Simulation
}

func (c *SimulationClient) mutate(ctx context.Context, m *SimulationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SimulationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SimulationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SimulationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SimulationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Simulation mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Simulation []ent.Hook
	}
	inters struct {
		Simulation []ent.Interceptor
	}
)