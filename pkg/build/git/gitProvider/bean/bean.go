/*
 * Copyright (c) 2024. Devtron Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bean

import (
	"github.com/devtron-labs/devtron/api/bean"
	"github.com/devtron-labs/devtron/internal/sql/constants"
)

type GitRegistry struct {
	Id                    int                `json:"id,omitempty" validate:"number"`
	Name                  string             `json:"name,omitempty" validate:"required"`
	Url                   string             `json:"url,omitempty"`
	UserName              string             `json:"userName,omitempty"`
	Password              string             `json:"password,omitempty"`
	SshPrivateKey         string             `json:"sshPrivateKey,omitempty"`
	AccessToken           string             `json:"accessToken,omitempty"`
	AuthMode              constants.AuthMode `json:"authMode,omitempty" validate:"required"`
	Active                bool               `json:"active"`
	UserId                int32              `json:"-"`
	GitHostId             int                `json:"gitHostId"`
	EnableTLSVerification bool               `json:"enableTLSVerification"`
	TLSConfig             bean.TLSConfig     `json:"tlsConfig"`
	IsCADataPresent       bool               `json:"isCADataPresent"`
	IsTLSCertDataPresent  bool               `json:"isTLSCertDataPresent"`
	IsTLSKeyDataPresent   bool               `json:"isTLSKeyDataPresent"`
}
