package content

import (
	"strings"
)

func GenerateActionsContent(name string) string {

	lowerCaseName := strings.ToLower(name)

	template := `import { createAsyncThunk } from '@reduxjs/toolkit';
import { instance } from '@/lib/axios';

export const create{pageName} = createAsyncThunk(
  '{lowerCaseName}/create{pageName}',
  async ({lowerCaseName}Data, { rejectWithValue }) => {
    try {
      const response = await instance.post('/{lowerCaseName}', {lowerCaseName}Data);
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response?.data || 'An error occurred');
    }
  }
);

export const get{pageName}s = createAsyncThunk('{lowerCaseName}/get{pageName}s', async () => {
  try {
    const response = await instance.get('/{lowerCaseName}');
    console.log(response.data);
    return response.data;
  } catch (error) {
    return 'error';
  }
});

export const get{pageName} = createAsyncThunk(
  '{lowerCaseName}/get{pageName}',
  async ({lowerCaseName}Id: string, { rejectWithValue }) => {
    try {
      const response = await instance.get('/{pageName}/${{lowerCaseName}Id}');
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response?.data || 'An error occurred');
    }
  }
);

export const update{pageName} = createAsyncThunk(
  '{lowerCaseName}/update{pageName}',
  async ({{lowerCaseName}Id, {lowerCaseName}Data } : any, { rejectWithValue }) => {
    try {
      const response = await instance.put('/{lowerCaseName}/${{lowerCaseName}Id}', {lowerCaseName}Data);
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response.data);
    }
  }
);

export const delete{pageName} = createAsyncThunk(
  '{lowerCaseName}/delete{pageName}',
  async ({lowerCaseName}Id: string, { rejectWithValue }) => {
    try {
      const response = await instance.delete('/{lowerCaseName}/${{lowerCaseName}Id}');
       return response.data;
    } catch (error) {
      return rejectWithValue(error.response?.data || 'An error occurred');
    }
  }
);
`
	finalCopy := strings.ReplaceAll(template, "{pageName}", name)
	finalCopy = strings.ReplaceAll(finalCopy, "{lowerCaseName}", lowerCaseName)
	return finalCopy
}
