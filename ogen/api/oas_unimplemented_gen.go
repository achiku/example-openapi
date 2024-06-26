// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddPet implements addPet operation.
//
// Add a new pet to the store.
//
// POST /pet
func (UnimplementedHandler) AddPet(ctx context.Context, req *Pet) (r AddPetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUser implements createUser operation.
//
// This can only be done by the logged in user.
//
// POST /user
func (UnimplementedHandler) CreateUser(ctx context.Context, req *User) (r *CreateUserDef, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUsersWithArrayInput implements createUsersWithArrayInput operation.
//
// Creates list of users with given input array.
//
// POST /user/createWithArray
func (UnimplementedHandler) CreateUsersWithArrayInput(ctx context.Context, req []User) (r *CreateUsersWithArrayInputDef, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateUsersWithListInput implements createUsersWithListInput operation.
//
// Creates list of users with given input array.
//
// POST /user/createWithList
func (UnimplementedHandler) CreateUsersWithListInput(ctx context.Context, req []User) (r *CreateUsersWithListInputDef, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteOrder implements deleteOrder operation.
//
// For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will
// generate API errors.
//
// DELETE /store/order/{orderId}
func (UnimplementedHandler) DeleteOrder(ctx context.Context, params DeleteOrderParams) (r DeleteOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeletePet implements deletePet operation.
//
// Deletes a pet.
//
// DELETE /pet/{petId}
func (UnimplementedHandler) DeletePet(ctx context.Context, params DeletePetParams) error {
	return ht.ErrNotImplemented
}

// DeleteUser implements deleteUser operation.
//
// This can only be done by the logged in user.
//
// DELETE /user/{username}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FindPetsByStatus implements findPetsByStatus operation.
//
// Multiple status values can be provided with comma separated strings.
//
// GET /pet/findByStatus
func (UnimplementedHandler) FindPetsByStatus(ctx context.Context, params FindPetsByStatusParams) (r FindPetsByStatusRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FindPetsByTags implements findPetsByTags operation.
//
// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
//
// Deprecated: schema marks this operation as deprecated.
//
// GET /pet/findByTags
func (UnimplementedHandler) FindPetsByTags(ctx context.Context, params FindPetsByTagsParams) (r FindPetsByTagsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetInventory implements getInventory operation.
//
// Returns a map of status codes to quantities.
//
// GET /store/inventory
func (UnimplementedHandler) GetInventory(ctx context.Context) (r GetInventoryOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetOrderById implements getOrderById operation.
//
// For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
//
// GET /store/order/{orderId}
func (UnimplementedHandler) GetOrderById(ctx context.Context, params GetOrderByIdParams) (r GetOrderByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPetById implements getPetById operation.
//
// Returns a single pet.
//
// GET /pet/{petId}
func (UnimplementedHandler) GetPetById(ctx context.Context, params GetPetByIdParams) (r GetPetByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserByName implements getUserByName operation.
//
// Get user by user name.
//
// GET /user/{username}
func (UnimplementedHandler) GetUserByName(ctx context.Context, params GetUserByNameParams) (r GetUserByNameRes, _ error) {
	return r, ht.ErrNotImplemented
}

// LoginUser implements loginUser operation.
//
// Logs user into the system.
//
// GET /user/login
func (UnimplementedHandler) LoginUser(ctx context.Context, params LoginUserParams) (r LoginUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// LogoutUser implements logoutUser operation.
//
// Logs out current logged in user session.
//
// GET /user/logout
func (UnimplementedHandler) LogoutUser(ctx context.Context) (r *LogoutUserDef, _ error) {
	return r, ht.ErrNotImplemented
}

// PlaceOrder implements placeOrder operation.
//
// Place an order for a pet.
//
// POST /store/order
func (UnimplementedHandler) PlaceOrder(ctx context.Context, req *Order) (r PlaceOrderRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePet implements updatePet operation.
//
// Update an existing pet.
//
// PUT /pet
func (UnimplementedHandler) UpdatePet(ctx context.Context, req *Pet) (r UpdatePetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePetWithForm implements updatePetWithForm operation.
//
// Updates a pet in the store with form data.
//
// POST /pet/{petId}
func (UnimplementedHandler) UpdatePetWithForm(ctx context.Context, req OptUpdatePetWithFormReq, params UpdatePetWithFormParams) error {
	return ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// This can only be done by the logged in user.
//
// PUT /user/{username}
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *User, params UpdateUserParams) (r UpdateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UploadFile implements uploadFile operation.
//
// Uploads an image.
//
// POST /pet/{petId}/uploadImage
func (UnimplementedHandler) UploadFile(ctx context.Context, req OptUploadFileReq, params UploadFileParams) (r *ApiResponse, _ error) {
	return r, ht.ErrNotImplemented
}
