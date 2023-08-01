import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

export const fetchOrders = createAsyncThunk(
  'orders/fetchOrders', 
  async () => {
    const response = await axios.get('http://localhost:8888/orders');
    return response.data;
  }
);

export const createOrder = createAsyncThunk(
  'orders/createOrder', 
  async (order, { dispatch }) => {
    let data = new URLSearchParams();
    data.append('PatientId', order.PatientId);
    data.append('Message', order.Message);
    const response = await axios.post('http://localhost:8888/orders/create', data);
    dispatch(fetchOrders());
    return response.data;
  }
);

export const editOrder = createAsyncThunk(
  'orders/editOrder', 
  async (updatedOrder, { dispatch }) => {
    let data = new URLSearchParams();
    data.append('OrderId', updatedOrder.Id);
    data.append('ModifiedMessage', updatedOrder.ModifiedMessage);
    const response = await axios.put('http://localhost:8888/orders/edit', data);
    dispatch(fetchOrders());
    return response.data;
  }
);

const initialState = {
  entities: [],
  loading: 'idle',
  error: null,
};

const ordersSlice = createSlice({
  name: 'orders',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchOrders.pending, (state) => {
        state.loading = 'loading';
      })
      .addCase(fetchOrders.fulfilled, (state, action) => {
        state.loading = 'succeeded';
        state.entities = action.payload;
      })
      .addCase(fetchOrders.rejected, (state, action) => {
      state.loading = 'failed';
      state.error = action.error.message;
    })
    .addCase(createOrder.fulfilled, (state, action) => {
      state.entities.push(action.payload);
    })
    .addCase(editOrder.fulfilled, (state, action) => {
        const index = state.entities.findIndex(order => order.Id === action.payload.Id);
        if (index !== -1) {
          state.entities[index] = action.payload;
        }
    });
  },
});

export default ordersSlice.reducer;

export const selectAllOrders = (state) => state.orders.entities;