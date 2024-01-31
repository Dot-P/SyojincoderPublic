import React from 'react';
import Caution from './Caution'
import AlgoTable from './AlgoTable';
import Searchbar from './SearchBar';

const AlgorithmPage: React.FC = () => {

  return (
    <div>
      <div className="w-full pt-10 px-4 sm:px-6 md:px-8 lg:ps-72">
        <h1 className="text-3xl text-black font-medium">Algorithm</h1>
        <br></br>
        <Searchbar></Searchbar>
        <br></br>
        <Caution></Caution>
      </div>
    </div>
  );
};

export default AlgorithmPage;
