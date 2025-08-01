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
 
import { uniq } from 'lodash-es';
import {
  type WorkflowNodeRegistry,
  type FlowNodeEntity,
} from '@flowgram-adapter/free-layout-editor';

/**
 * 获取有子节点的节点列表
 * @param nodes
 * @returns
 */
export const getNodesWithSubCanvas = (nodes: FlowNodeEntity[]) =>
  uniq(
    nodes
      .map(node => {
        const registry = node.getNodeRegistry() as WorkflowNodeRegistry;
        const subCanvas = registry?.meta?.subCanvas;

        return [
          node,
          // 子画布对应的所有子节点
          ...(subCanvas?.(node)?.canvasNode?.allCollapsedChildren || []),
        ];
      })
      .flat(),
  );
