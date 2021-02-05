//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBlockUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.BlockUser(1)
	if err != nil {
		t.Errorf("Users.BlockUser returned error: %v", err)
	}
}

func TestBlockUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.BlockUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\nGot: %+v", ErrUserNotFound, err)
	}
}

func TestBlockUser_BlockPrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.BlockUser(1)
	if !errors.Is(err, ErrUserBlockPrevented) {
		t.Errorf("Users.BlockUser error.\nExpected: %+v\nGot: %+v", ErrUserBlockPrevented, err)
	}
}

func TestBlockUser_UnknownError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/block", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.BlockUser(1)
	if err.Error() != want {
		t.Errorf("Users.BlockUser error.\nExpected: %s\nGot: %v", want, err)
	}
}

//  ------------------------  Unblock user -------------------------
func TestUnblockUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.UnblockUser(1)
	if err != nil {
		t.Errorf("Users.UnblockUser returned error: %v", err)
	}
}

func TestUnblockUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.UnblockUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.UnblockUser error.\nExpected: %v\nGot: %v", ErrUserNotFound, err)
	}
}

func TestUnblockUser_UnblockPrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.UnblockUser(1)
	if !errors.Is(err, ErrUserUnblockPrevented) {
		t.Errorf("Users.UnblockUser error.\nExpected: %v\nGot: %v", ErrUserUnblockPrevented, err)
	}
}

func TestUnblockUser_UnknownError(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/unblock", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusTeapot)
	})

	want := fmt.Sprintf("Received unexpected result code: %d", http.StatusTeapot)

	err := client.Users.UnblockUser(1)
	if err.Error() != want {
		t.Errorf("Users.UnblockUser error.\nExpected: %s\n\tGot: %v", want, err)
	}
}

func TestDeactivateUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.DeactivateUser(1)
	if err != nil {
		t.Errorf("Users.DeactivateUser returned error: %v", err)
	}
}

func TestDeactivateUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.DeactivateUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.DeactivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserNotFound, err)
	}
}

func TestDeactivateUser_DeactivatePrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/deactivate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.DeactivateUser(1)
	if !errors.Is(err, ErrUserDeactivatePrevented) {
		t.Errorf("Users.DeactivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserDeactivatePrevented, err)
	}
}

func TestActivateUser(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
	})

	err := client.Users.ActivateUser(1)
	if err != nil {
		t.Errorf("Users.ActivateUser returned error: %v", err)
	}
}

func TestActivateUser_ActivatePrevented(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
	})

	err := client.Users.ActivateUser(1)
	if !errors.Is(err, ErrUserActivatePrevented) {
		t.Errorf("Users.ActivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserActivatePrevented, err)
	}
}

func TestActivateUser_UserNotFound(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/activate", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusNotFound)
	})

	err := client.Users.ActivateUser(1)
	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("Users.ActivateUser error.\nExpected: %+v\n\tGot: %+v", ErrUserNotFound, err)
	}
}

func TestGetMemberships(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	path := fmt.Sprintf("/%susers/1/memberships", apiVersionPath)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_user_memberships.json")
	})

	opt := new(GetUserMembershipOptions)

	memberships, _, err := client.Users.GetUserMemberships(1, opt)
	require.NoError(t, err)

	want := []*UserMembership{{SourceID: 1, SourceName: "Project one", SourceType: "Project", AccessLevel: 20}, {SourceID: 3, SourceName: "Group three", SourceType: "Namespace", AccessLevel: 20}}
	assert.Equal(t, want, memberships)
}
