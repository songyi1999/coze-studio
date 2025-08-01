/*
 * Copyright 2025 coze-dev Authors
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
 
import { devtools, subscribeWithSelector } from 'zustand/middleware';
import { create } from 'zustand';
import { isFunction } from 'lodash-es';
import { produce } from 'immer';
import { type ShortCutStruct } from '@coze-agent-ide/tool-config';
import {
  type HookInfo,
  type LayoutInfo,
  type BackgroundImageInfo,
  type GetDraftBotInfoAgwData,
} from '@coze-arch/bot-api/playground_api';
import {
  type DefaultUserInputType,
  FileboxInfoMode,
  type PluginApi,
} from '@coze-arch/bot-api/developer_api';
import { botInputLengthService } from '@coze-agent-ide/bot-input-length-limit';

import {
  type SetterAction,
  setterActionFactory,
} from '../../utils/setter-factory';
import { getPluginApisFilterExample } from '../../utils/plugin-apis';
import {
  type VoicesInfo,
  type BotSuggestionConfig,
  type DatabaseInfo,
  type DatabaseList,
  type EnabledPluginApi,
  type ExtendOnboardingContent,
  type FileboxConfig,
  type KnowledgeConfig,
  type TaskManageInfo,
  type TimeCapsuleConfig,
  type TTSInfo,
  type VariableItem,
  type WorkFlowItemType,
} from '../../types/skill';
import { transformDto2Vo, transformVo2Dto } from './transform';
import {
  DEFAULT_BACKGROUND_IMAGE_LIST,
  DEFAULT_DATABASE,
  DEFAULT_KNOWLEDGE_CONFIG,
  DEFAULT_ONBOARDING_CONFIG,
  DEFAULT_SHORTCUT_CONFIG,
  DEFAULT_SUGGESTION_CONFIG,
  DEFAULT_TIME_CAPSULE_CONFIG,
  DEFAULT_TTS_CONFIG,
  DEFAULT_VOICES_INFO,
} from './defaults';

export const getDefaultBotSkillStore = (): BotSkillStore => ({
  pluginApis: [],
  workflows: [],
  knowledge: {
    dataSetList: [],
    dataSetInfo: DEFAULT_KNOWLEDGE_CONFIG(),
  },

  taskInfo: {
    user_task_allowed: false,
    data: [],
    task_list: [],
    loading: false,
  },
  variables: [],
  database: DEFAULT_DATABASE(),
  databaseList: [],
  onboardingContent: DEFAULT_ONBOARDING_CONFIG(),
  suggestionConfig: DEFAULT_SUGGESTION_CONFIG(),
  tts: DEFAULT_TTS_CONFIG(),
  voicesInfo: DEFAULT_VOICES_INFO(),
  timeCapsule: DEFAULT_TIME_CAPSULE_CONFIG(),
  filebox: {
    mode: FileboxInfoMode.Off,
  },
  backgroundImageInfoList: DEFAULT_BACKGROUND_IMAGE_LIST(),
  shortcut: DEFAULT_SHORTCUT_CONFIG(),
  layoutInfo: {},
  devHooks: {},
});

/** Persona & Prompt 区域 */
export interface BotSkillStore {
  // region Bot 和 Agent 维度共有 skills
  /** 已选的 plugin api */
  pluginApis: EnabledPluginApi[];
  /** 已选 workflow */
  workflows: WorkFlowItemType[];
  /** Knowledge 配置 */
  knowledge: KnowledgeConfig;
  // endregion

  // region Bot 维度独有 skills
  /**
   * task 配置
   *
   * 不含已添加的 task，已添加的在组件内独立管理
   */
  taskInfo: TaskManageInfo;
  /**
   * variable 默认值配置
   *
   * 不含右上角现值，现值为打开弹窗后请求获得的组件状态
   */
  variables: VariableItem[];
  /**
   * database 默认值配置
   *
   * 不含右上角现值，现值为打开弹窗后请求获得的组件状态
   */
  database: DatabaseInfo;
  /**
   * database 多表默认值配置
   *
   * 不含右上角现值，现值为打开弹窗后请求获得的组件状态
   */
  databaseList: DatabaseList;
  /** 开场白配置 */
  onboardingContent: ExtendOnboardingContent;
  /** 用户问题建议配置 */
  suggestionConfig: BotSuggestionConfig;
  // endregion
  /** 文字转语音 */
  tts: TTSInfo;
  /** 语音设置 上面 tts 在命名和含义划分上已经不准确了 但是牵扯甚广 */
  voicesInfo: VoicesInfo;
  // 时间胶囊
  timeCapsule: TimeCapsuleConfig;
  filebox: FileboxConfig;
  // 聊天背景图
  backgroundImageInfoList: BackgroundImageInfo[];
  // 快捷指令
  shortcut: ShortCutStruct;
  // hooks
  devHooks?: HookInfo;
  layoutInfo: LayoutInfo;
}

export interface BotSkillAction {
  setBotSkill: SetterAction<BotSkillStore>;
  setBotSkillByImmer: (update: (state: BotSkillStore) => void) => void;
  updateSkillPluginApis: (pluginApis: PluginApi[]) => void;
  updateSkillWorkflows: (workflows: WorkFlowItemType[]) => void;
  updateSkillKnowledgeDatasetList: (
    dataSetList: KnowledgeConfig['dataSetList'],
  ) => void;
  updateSkillKnowledgeDatasetInfo: (
    dataSetInfo: KnowledgeConfig['dataSetInfo'],
  ) => void;
  updateSkillTaskInfo: (taskInfo: Partial<TaskManageInfo>) => void;
  updateSkillDatabase: (database: Partial<DatabaseInfo>) => void;
  updateSkillDatabaseList: (database: DatabaseList) => void;
  updateSkillOnboarding: (
    onboarding:
      | Partial<ExtendOnboardingContent>
      | ((prev: ExtendOnboardingContent) => Partial<ExtendOnboardingContent>),
  ) => void;
  updateSkillLayoutInfo: (layoutInfo: LayoutInfo) => void;
  setBackgroundImageInfoList: (params: BackgroundImageInfo[]) => void;
  setSuggestionConfig: (config: Partial<BotSuggestionConfig>) => void;
  setDefaultUserInputType: (type: DefaultUserInputType) => void;
  transformDto2Vo: typeof transformDto2Vo;
  transformVo2Dto: typeof transformVo2Dto;
  initStore: (botData: GetDraftBotInfoAgwData) => void;
  clear: () => void;
}

export const useBotSkillStore = create<BotSkillStore & BotSkillAction>()(
  devtools(
    subscribeWithSelector((set, get) => ({
      ...getDefaultBotSkillStore(),

      setBotSkill: setterActionFactory<BotSkillStore>(set),
      setBotSkillByImmer: update =>
        set(produce<BotSkillStore>(botSkill => update(botSkill))),
      updateSkillPluginApis: (pluginApis: PluginApi[]) => {
        set(s => ({
          ...s,
          pluginApis: getPluginApisFilterExample(pluginApis),
        }));
      },
      updateSkillWorkflows: workflows => set(s => ({ ...s, workflows })),
      updateSkillKnowledgeDatasetList: dataSetList =>
        set(
          produce<BotSkillStore>(s => {
            s.knowledge.dataSetList = dataSetList;
          }),
        ),
      updateSkillKnowledgeDatasetInfo: dataSetInfo =>
        set(
          produce<BotSkillStore>(s => {
            s.knowledge.dataSetInfo = dataSetInfo;
          }),
        ),
      updateSkillTaskInfo: taskInfo =>
        set(s => ({
          ...s,
          taskInfo: { ...s.taskInfo, ...taskInfo },
        })),
      updateSkillDatabase: database =>
        set(s => ({
          ...s,
          database: { ...s.database, ...database },
        })),
      updateSkillDatabaseList: databaseList =>
        set(
          produce<BotSkillStore>(s => {
            s.databaseList = databaseList;
          }),
        ),
      updateSkillOnboarding: update =>
        set(s => ({
          ...s,
          onboardingContent: {
            ...s.onboardingContent,
            ...(isFunction(update) ? update(s.onboardingContent) : update),
          },
        })),
      updateSkillLayoutInfo: layoutInfo => {
        set(s => ({
          ...s,
          layoutInfo,
        }));
      },
      setSuggestionConfig: config =>
        set(s => ({
          ...s,
          suggestionConfig: { ...s.suggestionConfig, ...config },
        })),
      setBackgroundImageInfoList: config =>
        set(s => ({
          ...s,
          backgroundImageInfoList: [...config],
        })),
      setDefaultUserInputType: inputType =>
        set(
          state =>
            produce(state, draft => {
              draft.voicesInfo.defaultUserInputType = inputType;
            }),
          false,
          'setDefaultUserInputType',
        ),
      transformDto2Vo,
      transformVo2Dto,
      initStore: botData => {
        const { bot_info: botInfo, bot_option_data: optionData } = botData;
        set({
          pluginApis: transformDto2Vo.plugin(
            botInfo?.plugin_info_list,
            optionData?.plugin_detail_map,
            optionData?.plugin_api_detail_map,
          ),
          workflows: transformDto2Vo.workflow(
            botInfo?.workflow_info_list,
            optionData?.workflow_detail_map,
          ),
          knowledge: transformDto2Vo.knowledge(
            botInfo?.knowledge,
            optionData?.knowledge_detail_map,
          ),
          taskInfo: transformDto2Vo.task(botInfo?.task_info),
          variables: transformDto2Vo.variables(botInfo?.variable_list),
          databaseList: transformDto2Vo.databaseList(botInfo?.database_list),
          timeCapsule: transformDto2Vo.timeCapsule(
            botInfo?.bot_tag_info?.time_capsule_info,
          ),
          filebox: transformDto2Vo.filebox(botInfo?.filebox_info),

          onboardingContent:
            botInputLengthService.sliceWorkInfoOnboardingByMaxLength(
              transformDto2Vo.onboarding(botInfo?.onboarding_info),
            ),
          suggestionConfig: transformDto2Vo.suggestionConfig(
            botInfo?.suggest_reply_info,
          ),
          tts: transformDto2Vo.tts(botInfo?.voices_info),
          voicesInfo: transformDto2Vo.voicesInfo(botInfo.voices_info),
          backgroundImageInfoList: botInfo?.background_image_info_list ?? [],
          shortcut: transformDto2Vo.shortcut(
            botInfo?.shortcut_sort ?? [],
            optionData?.shortcut_command_list,
          ),
          devHooks: transformDto2Vo.hookInfo(botInfo?.hook_info),
          layoutInfo: transformDto2Vo.layoutInfo(botInfo?.layout_info),
        });
      },
      clear: () => {
        set({ ...getDefaultBotSkillStore() });
      },
    })),
    {
      enabled: IS_DEV_MODE,
      name: 'botStudio.botDetail.botSkill',
    },
  ),
);
