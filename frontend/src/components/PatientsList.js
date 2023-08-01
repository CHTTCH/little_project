import React, { useEffect, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { fetchPatients, selectAllPatients } from '../redux/patientsSlice';
import { fetchOrders, selectAllOrders, createOrder, editOrder } from '../redux/ordersSlice';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemButton from '@mui/material/ListItemButton';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import IconButton from '@mui/material/IconButton';
import EditIcon from '@mui/icons-material/Edit';
import AddIcon from '@mui/icons-material/Add';
import SaveIcon from '@mui/icons-material/Save';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';

const PatientsList = () => {
  const dispatch = useDispatch();
  const patients = useSelector(selectAllPatients);
  const orders = useSelector(selectAllOrders);
  const patientsStatus = useSelector((state) => state.patients.status);
  const ordersStatus = useSelector((state) => state.orders.loading);
  const error = useSelector((state) => state.patients.error);

  const [open, setOpen] = useState(false);
  const [selectedOrder, setSelectedOrder] = useState(null);
  const [selectedPatient, setSelectedPatient] = useState(null);
  const [editMode, setEditMode] = useState(false);
  const [message, setMessage] = useState('');

  useEffect(() => {
    if (patientsStatus === 'idle') {
      dispatch(fetchPatients());
    }
    if (ordersStatus === 'idle') {
      dispatch(fetchOrders());
    }
  }, [patientsStatus, ordersStatus, dispatch]);
  
  const handleListItemClick = (patient) => {
    console.log(`You clicked patient with id ${patient.Id}`);
    const order = orders.find((order) => order.Id === patient.OrderId);
    setSelectedOrder(order);
    setSelectedPatient(patient);
    setOpen(true);
    if (order) {
      setEditMode(false);
      setMessage(order.Message);
    }
  };
  
  const handleClose = () => {
    setOpen(false);
    setEditMode(false);
  };

  const handleAddClick = () => {
    setEditMode(true);
    setMessage('');
  };

  const handleEditClick = () => {
    setEditMode(true);
  };

  const handleSaveClick = async () => {
    if (editMode) {
      if (selectedOrder) {
        await dispatch(editOrder({ Id: selectedOrder.Id, ModifiedMessage: message }))
          .then(() => {
            dispatch(fetchPatients());
          });

        const updatedOrders = await dispatch(fetchOrders());
        const updatedOrder = updatedOrders.payload.find((order) => order.Id === selectedPatient.OrderId);
        setSelectedOrder(updatedOrder);
        setMessage(updatedOrder ? updatedOrder.Message : '')
      } else {
        await dispatch(createOrder({ PatientId: selectedPatient.Id, Message: message }))
          .then(() => {
            dispatch(fetchPatients());
          });

        const updatedOrders = await dispatch(fetchOrders());
        const updatedPatients = await dispatch(fetchPatients());
        const updatedSelectedPatient = updatedPatients.payload.find((patient) => patient.Id === selectedPatient.Id);
        const updatedOrder = updatedOrders.payload.find((order) => order.Id === updatedSelectedPatient.OrderId);
        setSelectedOrder(updatedOrder);
        setSelectedPatient(updatedSelectedPatient);
        setMessage(updatedOrder ? updatedOrder.Message : '');
      }
    }
    setEditMode(false);
  };
  

  const handleChange = (event) => {
    setMessage(event.target.value);
  };

  let content;

  if (patientsStatus === 'loading' || ordersStatus === 'loading') {
    content = <div>Loading...</div>;
  } else if (patientsStatus === 'succeeded') {
    content = patients.slice(0, 5).map((patient) => (
      <ListItem key={patient.Id}>
        <ListItemButton key={patient.Id} onClick={() => handleListItemClick(patient)}>
          <ListItemText primary={`(${patient.Id}) ${patient.Name}`} />
        </ListItemButton>
      </ListItem>
    ));
  } else if (patientsStatus === 'failed') {
    content = <div>{error}</div>;
  }

  return (
    <div>
      <List>
        {content}
      </List>
      <Dialog open={open} onClose={handleClose}>
        <DialogTitle>
              <Box display="flex" alignItems="center" justifyContent="space-between">
                <Typography variant="h6">醫囑</Typography>
                {!editMode && <IconButton onClick={selectedOrder ? handleEditClick : handleAddClick}>
                  {selectedOrder ? <EditIcon /> : <AddIcon />}
                </IconButton>}
                {editMode && <IconButton onClick={handleSaveClick}><SaveIcon /></IconButton>}
              </Box>
        </DialogTitle>
        <DialogContent>
          {!editMode ? (selectedOrder ? selectedOrder.Message : '該病患目前沒有醫囑') : 
              <TextField
                autoFocus
                margin="dense"
                id="message"
                label="訊息"
                type="text"
                fullWidth
                value={message}
                onChange={handleChange}
              />
            }
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default PatientsList;
