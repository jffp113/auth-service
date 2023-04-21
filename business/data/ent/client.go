// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"com.cross-join.crossviewer.authservice/business/data/ent/migrate"

	"com.cross-join.crossviewer.authservice/business/data/ent/claim"
	"com.cross-join.crossviewer.authservice/business/data/ent/group"
	"com.cross-join.crossviewer.authservice/business/data/ent/role"
	"com.cross-join.crossviewer.authservice/business/data/ent/user"
	"com.cross-join.crossviewer.authservice/business/data/ent/usersgroups"
	"com.cross-join.crossviewer.authservice/business/data/ent/usersroles"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Claim is the client for interacting with the Claim builders.
	Claim *ClaimClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UsersGroups is the client for interacting with the UsersGroups builders.
	UsersGroups *UsersGroupsClient
	// UsersRoles is the client for interacting with the UsersRoles builders.
	UsersRoles *UsersRolesClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Claim = NewClaimClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
	c.UsersGroups = NewUsersGroupsClient(c.config)
	c.UsersRoles = NewUsersRolesClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Claim:       NewClaimClient(cfg),
		Group:       NewGroupClient(cfg),
		Role:        NewRoleClient(cfg),
		User:        NewUserClient(cfg),
		UsersGroups: NewUsersGroupsClient(cfg),
		UsersRoles:  NewUsersRolesClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Claim:       NewClaimClient(cfg),
		Group:       NewGroupClient(cfg),
		Role:        NewRoleClient(cfg),
		User:        NewUserClient(cfg),
		UsersGroups: NewUsersGroupsClient(cfg),
		UsersRoles:  NewUsersRolesClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Claim.
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
	for _, n := range []interface{ Use(...Hook) }{
		c.Claim, c.Group, c.Role, c.User, c.UsersGroups, c.UsersRoles,
	} {
		n.Use(hooks...)
	}
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	for _, n := range []interface{ Intercept(...Interceptor) }{
		c.Claim, c.Group, c.Role, c.User, c.UsersGroups, c.UsersRoles,
	} {
		n.Intercept(interceptors...)
	}
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ClaimMutation:
		return c.Claim.mutate(ctx, m)
	case *GroupMutation:
		return c.Group.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *UsersGroupsMutation:
		return c.UsersGroups.mutate(ctx, m)
	case *UsersRolesMutation:
		return c.UsersRoles.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ClaimClient is a client for the Claim schema.
type ClaimClient struct {
	config
}

// NewClaimClient returns a client for the Claim from the given config.
func NewClaimClient(c config) *ClaimClient {
	return &ClaimClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `claim.Hooks(f(g(h())))`.
func (c *ClaimClient) Use(hooks ...Hook) {
	c.hooks.Claim = append(c.hooks.Claim, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `claim.Intercept(f(g(h())))`.
func (c *ClaimClient) Intercept(interceptors ...Interceptor) {
	c.inters.Claim = append(c.inters.Claim, interceptors...)
}

// Create returns a builder for creating a Claim entity.
func (c *ClaimClient) Create() *ClaimCreate {
	mutation := newClaimMutation(c.config, OpCreate)
	return &ClaimCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Claim entities.
func (c *ClaimClient) CreateBulk(builders ...*ClaimCreate) *ClaimCreateBulk {
	return &ClaimCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Claim.
func (c *ClaimClient) Update() *ClaimUpdate {
	mutation := newClaimMutation(c.config, OpUpdate)
	return &ClaimUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClaimClient) UpdateOne(cl *Claim) *ClaimUpdateOne {
	mutation := newClaimMutation(c.config, OpUpdateOne, withClaim(cl))
	return &ClaimUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClaimClient) UpdateOneID(id int) *ClaimUpdateOne {
	mutation := newClaimMutation(c.config, OpUpdateOne, withClaimID(id))
	return &ClaimUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Claim.
func (c *ClaimClient) Delete() *ClaimDelete {
	mutation := newClaimMutation(c.config, OpDelete)
	return &ClaimDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ClaimClient) DeleteOne(cl *Claim) *ClaimDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ClaimClient) DeleteOneID(id int) *ClaimDeleteOne {
	builder := c.Delete().Where(claim.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClaimDeleteOne{builder}
}

// Query returns a query builder for Claim.
func (c *ClaimClient) Query() *ClaimQuery {
	return &ClaimQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeClaim},
		inters: c.Interceptors(),
	}
}

// Get returns a Claim entity by its id.
func (c *ClaimClient) Get(ctx context.Context, id int) (*Claim, error) {
	return c.Query().Where(claim.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClaimClient) GetX(ctx context.Context, id int) *Claim {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Claim.
func (c *ClaimClient) QueryUser(cl *Claim) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(claim.Table, claim.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, claim.UserTable, claim.UserColumn),
		)
		fromV = sqlgraph.Neighbors(cl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ClaimClient) Hooks() []Hook {
	return c.hooks.Claim
}

// Interceptors returns the client interceptors.
func (c *ClaimClient) Interceptors() []Interceptor {
	return c.inters.Claim
}

func (c *ClaimClient) mutate(ctx context.Context, m *ClaimMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ClaimCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ClaimUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ClaimUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ClaimDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Claim mutation op: %q", m.Op())
	}
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `group.Intercept(f(g(h())))`.
func (c *GroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.Group = append(c.inters.Group, interceptors...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthUsersGroups queries the auth_users_groups edge of a Group.
func (c *GroupClient) QueryAuthUsersGroups(gr *Group) *UsersGroupsQuery {
	query := (&UsersGroupsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(usersgroups.Table, usersgroups.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, group.AuthUsersGroupsTable, group.AuthUsersGroupsColumn),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// Interceptors returns the client interceptors.
func (c *GroupClient) Interceptors() []Interceptor {
	return c.inters.Group
}

func (c *GroupClient) mutate(ctx context.Context, m *GroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Group mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id int) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id int) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id int) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id int) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Role.
func (c *RoleClient) QueryUsers(r *Role) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.UsersTable, role.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthUsersRoles queries the auth_users_roles edge of a Role.
func (c *RoleClient) QueryAuthUsersRoles(r *Role) *UsersRolesQuery {
	query := (&UsersRolesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(usersroles.Table, usersroles.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, role.AuthUsersRolesTable, role.AuthUsersRolesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a User.
func (c *UserClient) QueryRoles(u *User) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.RolesTable, user.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryClaims queries the claims edge of a User.
func (c *UserClient) QueryClaims(u *User) *ClaimQuery {
	query := (&ClaimClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(claim.Table, claim.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.ClaimsTable, user.ClaimsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := (&GroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.GroupsTable, user.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthUsersRoles queries the auth_users_roles edge of a User.
func (c *UserClient) QueryAuthUsersRoles(u *User) *UsersRolesQuery {
	query := (&UsersRolesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usersroles.Table, usersroles.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.AuthUsersRolesTable, user.AuthUsersRolesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAuthUsersGroups queries the auth_users_groups edge of a User.
func (c *UserClient) QueryAuthUsersGroups(u *User) *UsersGroupsQuery {
	query := (&UsersGroupsClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(usersgroups.Table, usersgroups.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.AuthUsersGroupsTable, user.AuthUsersGroupsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// UsersGroupsClient is a client for the UsersGroups schema.
type UsersGroupsClient struct {
	config
}

// NewUsersGroupsClient returns a client for the UsersGroups from the given config.
func NewUsersGroupsClient(c config) *UsersGroupsClient {
	return &UsersGroupsClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usersgroups.Hooks(f(g(h())))`.
func (c *UsersGroupsClient) Use(hooks ...Hook) {
	c.hooks.UsersGroups = append(c.hooks.UsersGroups, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `usersgroups.Intercept(f(g(h())))`.
func (c *UsersGroupsClient) Intercept(interceptors ...Interceptor) {
	c.inters.UsersGroups = append(c.inters.UsersGroups, interceptors...)
}

// Create returns a builder for creating a UsersGroups entity.
func (c *UsersGroupsClient) Create() *UsersGroupsCreate {
	mutation := newUsersGroupsMutation(c.config, OpCreate)
	return &UsersGroupsCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UsersGroups entities.
func (c *UsersGroupsClient) CreateBulk(builders ...*UsersGroupsCreate) *UsersGroupsCreateBulk {
	return &UsersGroupsCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UsersGroups.
func (c *UsersGroupsClient) Update() *UsersGroupsUpdate {
	mutation := newUsersGroupsMutation(c.config, OpUpdate)
	return &UsersGroupsUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersGroupsClient) UpdateOne(ug *UsersGroups) *UsersGroupsUpdateOne {
	mutation := newUsersGroupsMutation(c.config, OpUpdateOne, withUsersGroups(ug))
	return &UsersGroupsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersGroupsClient) UpdateOneID(id int) *UsersGroupsUpdateOne {
	mutation := newUsersGroupsMutation(c.config, OpUpdateOne, withUsersGroupsID(id))
	return &UsersGroupsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UsersGroups.
func (c *UsersGroupsClient) Delete() *UsersGroupsDelete {
	mutation := newUsersGroupsMutation(c.config, OpDelete)
	return &UsersGroupsDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersGroupsClient) DeleteOne(ug *UsersGroups) *UsersGroupsDeleteOne {
	return c.DeleteOneID(ug.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersGroupsClient) DeleteOneID(id int) *UsersGroupsDeleteOne {
	builder := c.Delete().Where(usersgroups.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersGroupsDeleteOne{builder}
}

// Query returns a query builder for UsersGroups.
func (c *UsersGroupsClient) Query() *UsersGroupsQuery {
	return &UsersGroupsQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUsersGroups},
		inters: c.Interceptors(),
	}
}

// Get returns a UsersGroups entity by its id.
func (c *UsersGroupsClient) Get(ctx context.Context, id int) (*UsersGroups, error) {
	return c.Query().Where(usersgroups.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersGroupsClient) GetX(ctx context.Context, id int) *UsersGroups {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UsersGroups.
func (c *UsersGroupsClient) QueryUser(ug *UsersGroups) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ug.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersgroups.Table, usersgroups.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usersgroups.UserTable, usersgroups.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ug.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoles queries the roles edge of a UsersGroups.
func (c *UsersGroupsClient) QueryRoles(ug *UsersGroups) *GroupQuery {
	query := (&GroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ug.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersgroups.Table, usersgroups.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usersgroups.RolesTable, usersgroups.RolesColumn),
		)
		fromV = sqlgraph.Neighbors(ug.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UsersGroupsClient) Hooks() []Hook {
	return c.hooks.UsersGroups
}

// Interceptors returns the client interceptors.
func (c *UsersGroupsClient) Interceptors() []Interceptor {
	return c.inters.UsersGroups
}

func (c *UsersGroupsClient) mutate(ctx context.Context, m *UsersGroupsMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UsersGroupsCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UsersGroupsUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UsersGroupsUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UsersGroupsDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UsersGroups mutation op: %q", m.Op())
	}
}

// UsersRolesClient is a client for the UsersRoles schema.
type UsersRolesClient struct {
	config
}

// NewUsersRolesClient returns a client for the UsersRoles from the given config.
func NewUsersRolesClient(c config) *UsersRolesClient {
	return &UsersRolesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `usersroles.Hooks(f(g(h())))`.
func (c *UsersRolesClient) Use(hooks ...Hook) {
	c.hooks.UsersRoles = append(c.hooks.UsersRoles, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `usersroles.Intercept(f(g(h())))`.
func (c *UsersRolesClient) Intercept(interceptors ...Interceptor) {
	c.inters.UsersRoles = append(c.inters.UsersRoles, interceptors...)
}

// Create returns a builder for creating a UsersRoles entity.
func (c *UsersRolesClient) Create() *UsersRolesCreate {
	mutation := newUsersRolesMutation(c.config, OpCreate)
	return &UsersRolesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UsersRoles entities.
func (c *UsersRolesClient) CreateBulk(builders ...*UsersRolesCreate) *UsersRolesCreateBulk {
	return &UsersRolesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UsersRoles.
func (c *UsersRolesClient) Update() *UsersRolesUpdate {
	mutation := newUsersRolesMutation(c.config, OpUpdate)
	return &UsersRolesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersRolesClient) UpdateOne(ur *UsersRoles) *UsersRolesUpdateOne {
	mutation := newUsersRolesMutation(c.config, OpUpdateOne, withUsersRoles(ur))
	return &UsersRolesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersRolesClient) UpdateOneID(id int) *UsersRolesUpdateOne {
	mutation := newUsersRolesMutation(c.config, OpUpdateOne, withUsersRolesID(id))
	return &UsersRolesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UsersRoles.
func (c *UsersRolesClient) Delete() *UsersRolesDelete {
	mutation := newUsersRolesMutation(c.config, OpDelete)
	return &UsersRolesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersRolesClient) DeleteOne(ur *UsersRoles) *UsersRolesDeleteOne {
	return c.DeleteOneID(ur.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersRolesClient) DeleteOneID(id int) *UsersRolesDeleteOne {
	builder := c.Delete().Where(usersroles.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersRolesDeleteOne{builder}
}

// Query returns a query builder for UsersRoles.
func (c *UsersRolesClient) Query() *UsersRolesQuery {
	return &UsersRolesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUsersRoles},
		inters: c.Interceptors(),
	}
}

// Get returns a UsersRoles entity by its id.
func (c *UsersRolesClient) Get(ctx context.Context, id int) (*UsersRoles, error) {
	return c.Query().Where(usersroles.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersRolesClient) GetX(ctx context.Context, id int) *UsersRoles {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UsersRoles.
func (c *UsersRolesClient) QueryUser(ur *UsersRoles) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ur.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersroles.Table, usersroles.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usersroles.UserTable, usersroles.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ur.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRoles queries the roles edge of a UsersRoles.
func (c *UsersRolesClient) QueryRoles(ur *UsersRoles) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ur.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(usersroles.Table, usersroles.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usersroles.RolesTable, usersroles.RolesColumn),
		)
		fromV = sqlgraph.Neighbors(ur.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UsersRolesClient) Hooks() []Hook {
	return c.hooks.UsersRoles
}

// Interceptors returns the client interceptors.
func (c *UsersRolesClient) Interceptors() []Interceptor {
	return c.inters.UsersRoles
}

func (c *UsersRolesClient) mutate(ctx context.Context, m *UsersRolesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UsersRolesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UsersRolesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UsersRolesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UsersRolesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UsersRoles mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Claim, Group, Role, User, UsersGroups, UsersRoles []ent.Hook
	}
	inters struct {
		Claim, Group, Role, User, UsersGroups, UsersRoles []ent.Interceptor
	}
)
