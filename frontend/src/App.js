import React from 'react';
import { Provider } from 'react-redux';
import { store } from './redux/store';
import PatientsList from './components/PatientsList';

const App = () => (
  <Provider store={store}>
    <PatientsList />
  </Provider>
);

export default App;
