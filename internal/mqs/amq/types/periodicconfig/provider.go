// Copyright 2023 The Ryan SU Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package periodicconfig

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-job/ent"
	"github.com/suyuan32/simple-admin-job/ent/task"
	"github.com/suyuan32/simple-admin-job/internal/utils/dberrorhandler"
)

type EntConfigProvider struct {
	DB *ent.Client
}

func NewEntConfigProvider(db *ent.Client) *EntConfigProvider {
	return &EntConfigProvider{
		db,
	}
}

func (e *EntConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	configData, err := e.DB.Task.Query().Where(task.StatusEQ(common.StatusNormal)).All(context.Background())
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(logx.WithContext(context.Background()), err, "get task config error")
	}

	var result []*asynq.PeriodicTaskConfig

	for _, v := range configData {
		result = append(result, &asynq.PeriodicTaskConfig{
			Cronspec: v.CronExpression,
			Task:     asynq.NewTask(v.Pattern, []byte(v.Payload)),
			Opts:     nil,
		})
	}

	return result, nil
}
