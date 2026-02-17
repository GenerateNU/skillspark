import { defineConfig } from 'orval';
import path from 'path';

export default defineConfig({
  api: {
    output: {
      mode: 'tags-split',
      target: './src/generated',
      client: 'react-query',
      mock: false,
      clean: true,
      prettier: true,
      indexFiles: true,
      override: {
        mutator: {
          path: path.resolve(__dirname, 'src/apiClient.ts'),
          name: 'customInstance',
        },
        query: {
          useQuery: true,
          useMutation: true,
          signal: true,
          version: 5,
        },
      },
    },
    input: {
      target: '../../../backend/api/openapi.yaml',
    },
  },
});