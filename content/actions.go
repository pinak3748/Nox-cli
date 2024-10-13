package content

import (
	"strings"
)

func GenerateActionsContent(name string) string {
	template := `import { createAsyncThunk } from '@reduxjs/toolkit';
import { {pageName}Item } from '@/constants/types';
import { instance } from '@/lib/axios';

export const create{pageName} = createAsyncThunk(
  '{pageName}/create{pageName}',
  async ({pageName}Data: Omit<{pageName}Item, 'id'>, { rejectWithValue }) => {
    try {
      const response = await instance.post('/{pageName}', {pageName}Data);
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response?.data || 'An error occurred');
    }
  }
);

export const get{pageName}s = createAsyncThunk('{pageName}/get{pageName}s', async () => {
  try {
    const response = await instance.get('/{pageName}');
    console.log(response.data);
    return response.data;
  } catch (error) {
    return 'error';
  }
});

export const get{pageName} = createAsyncThunk(
  '{pageName}/get{pageName}',
  async ({pageName}Id: string, { rejectWithValue }) => {
    try {
      const response = await instance.get('/{pageName}/${{pageName}Id}');
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response?.data || 'An error occurred');
    }
  }
);

export const update{pageName} = createAsyncThunk(
  '{pageName}/update{pageName}',
  async ({pageName}Data: {pageName}Item, { rejectWithValue }) => {
    try {
      const response = await instance.put('/{pageName}/' + ${pageName}Data.id, {pageName}Data);
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response.data);
    }
  }
);

export const delete{pageName} = createAsyncThunk(
  '{pageName}/delete{pageName}',
  async ({pageName}Id: string, { rejectWithValue }) => {
    try {
      await instance.delete('/{pageName}/${{pageName}Id}');
      return {pageName}Id;
    } catch (error) {
      return rejectWithValue(error.response?.data || 'An error occurred');
    }
  }
);
`

	return strings.ReplaceAll(template, "{pageName}", name)
}
