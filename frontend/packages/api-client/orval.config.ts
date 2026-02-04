import { defineConfig } from 'orval';

export default defineConfig({
  api: {
    input: '../../../../backend/api/openapi.yaml', // Point to backend spec
    output: {
      mode: 'tags-split',
      target: './src/generated',
      client: 'react-query',
      override: {
        mutator: {
          path: './src/mutator.ts',
          name: 'customInstance',
        },
      },
    },
  },
});