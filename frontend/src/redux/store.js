import { configureStore } from '@reduxjs/toolkit';
import patientsReducer from './patientsSlice';
import ordersReducer from './ordersSlice';

export const store = configureStore({
  reducer: {
    patients: patientsReducer,
    orders: ordersReducer,
  },
});