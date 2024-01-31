import React, { useState } from 'react';
import Table from './Table';
import SearchBox from './Searchbox';
import Graph1 from './Graph1';
import Graph2 from './Graph2';

const DashBoard: React.FC = () => {

  const [userName, setUserName] = useState('');

  return (
    <div>
      <div className="w-full pt-10 px-4 sm:px-6 md:px-8 lg:ps-72">
        <h1 className="text-3xl text-black font-medium">Dashboard</h1>
        <br></br>
        <SearchBox setUserName={setUserName} />
        <div className="flex flex-row flex-wrap justify-center space-x-4">
          <Graph1 />
          <Graph2 />
        </div>
        <Table />
      </div>
    </div>
  );
};

export default DashBoard;
