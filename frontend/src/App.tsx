import React, { useEffect } from 'react';
import { Route, Routes, useLocation } from 'react-router-dom';
import DashBoard from './component/Dashbord';
import SideMenus from './component/Sidemenu';
import AlgorithmPage from './component/AlgorithmPage'; 
import { DataProvider } from './DataContext';
import './App.css';

function App() {
  const location = useLocation();

  useEffect(() => {
    require('preline/preline');
  }, []);

  useEffect(() => {
    // @ts-ignore
    HSStaticMethods.autoInit();
  }, [location.pathname]);

  return (
    <div className="App">
      <SideMenus />
      <DataProvider>
        <Routes>
          <Route path="/algorithm" element={<AlgorithmPage />} />
          <Route path="/" element={<DashBoard />} />
        </Routes>
      </DataProvider>
    </div>
  );
}

export default App;
