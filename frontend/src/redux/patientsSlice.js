import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

export const fetchPatients = createAsyncThunk(
    'patients/fetchPatients',
    async () => {
      const response = await axios.get('http://localhost:8888/patients');
      return response.data;
    }
);

const initialState = {
    patients: [],
    status: 'idle',
    error: null,
};

export const patientsSlice = createSlice({
    name: 'patients',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
      builder
        .addCase(fetchPatients.pending, (state) => {
          state.status = 'loading';
        })
        .addCase(fetchPatients.fulfilled, (state, action) => {
          state.status = 'succeeded';
          state.patients = action.payload;
        })
        .addCase(fetchPatients.rejected, (state, action) => {
          state.status = 'failed';
          state.error = action.error.message;
        });
    },
});

export const selectAllPatients = (state) => state.patients.patients;

export default patientsSlice.reducer;