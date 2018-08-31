/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/okta"

	"github.com/okta/okta-sdk-golang/tests"
)

func Test_can_create_a_bookmark_applicaiton(t *testing.T) {
	client := tests.NewClient()

	bookmarkApplicationSettingsApplication := okta.BookmarkApplicationSettingsApplication{
		RequestIntegration: new(bool),
		Url:                "https://example.com/bookmark.htm",
	}

	bookmarkApplicationSettings := okta.BookmarkApplicationSettings{
		App: &bookmarkApplicationSettingsApplication,
	}

	bookmarkApplication := okta.NewBookmarkApplication()
	bookmarkApplication.Settings = &bookmarkApplicationSettings

	application, _, err := client.Application.CreateApplication(bookmarkApplication, nil)
	require.NoError(t, err, "Creating an application should not error")

	assert.IsType(t, &okta.BookmarkApplication{}, application, "Application type returned was incorrect")

	client.Application.DeactivateApplication(application.(*okta.BookmarkApplication).Id)
	_, err = client.Application.DeleteApplication(application.(*okta.BookmarkApplication).Id)

	require.NoError(t, err, "Deleting an application should not error")
}