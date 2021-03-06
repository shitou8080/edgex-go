/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package interfaces

import (
	"github.com/edgexfoundry/edgex-go/internal/support/logging/filter"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type Persistence interface {
	Add(logEntry models.LogEntry) error
	Remove(criteria filter.Criteria) (int, error)
	Find(criteria filter.Criteria) ([]models.LogEntry, error)
	CloseSession()
}
