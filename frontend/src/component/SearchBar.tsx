import React, { useState } from 'react';
import './SearchBar.css'; 

const Searchbar: React.FC = () => {

  interface Problem {
    problem_id: string;
    difficulties: string;
    classifier: string;
    url: string;
  }
  
  const [selectedOption, setSelectedOption] = useState('');
  const [problems, setProblems] = useState<Problem[]>([]);

  const handleOptionChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedOption(e.target.value);
    console.log("Selected option:", e.target.value);
  };

  const handleSaveChanges = async () => {
    console.log("Sending option:", selectedOption); 
    try {
      // APIリクエストを送信
      const response = await fetch("https://syojincoder-3d33c5ad4332.herokuapp.com/algo_search", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ option: selectedOption }),
      });

      if (!response.ok) {
        throw new Error('Something went wrong');
      }

      // レスポンスの処理
      const data = await response.json();
      setProblems(data);
      console.log("Response data:", data); 
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
<div>
<button type="button" className="py-3 px-4 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none" data-hs-overlay="#hs-toggle-password-modal-example">
  Select Algorithm
</button>

<br></br>

<div id="hs-toggle-password-modal-example" className="hs-overlay hidden w-full h-full fixed top-0 start-0 z-[60] overflow-x-hidden overflow-y-auto pointer-events-none">
  <div className="hs-overlay-open:mt-7 hs-overlay-open:opacity-100 hs-overlay-open:duration-500 mt-0 opacity-0 ease-out transition-all sm:max-w-lg sm:w-full m-3 sm:mx-auto">
    <div className="flex flex-col bg-white border shadow-sm rounded-xl pointer-events-auto">
      <div className="flex justify-between items-center py-3 px-4 border-b">
        <h3 className="font-bold text-gray-800">
          検索したいアルゴリズムを選択してください。
        </h3>
        <button type="button" className="hs-dropdown-toggle inline-flex flex-shrink-0 justify-center items-center h-8 w-8 rounded-lg text-gray-500 hover:text-gray-400 focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 focus:ring-offset-white transition-all text-sm" data-hs-overlay="#hs-toggle-password-modal-example">
          <span className="sr-only">Close</span>
          <svg className="flex-shrink-0 w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
        </button>
      </div>
    <div className="p-4 min-h-[15rem] overflow-y-auto">

      <div className="relative">
        <select onChange={handleOptionChange} value={selectedOption} className="select-style">
        <option value="Name">Select the option ...</option>
        <option value="1">全探索</option>
        <option value="2">累積和及びimos法</option>
        <option value="3">二分探索</option>
        <option value="4">座標圧縮、ランレングス圧縮</option>
        <option value="5">動的計画法</option>
        <option value="6">尺取り法</option>
        <option value="7">BFS</option>
        <option value="15">DFS</option>
        <option value="8">Union-Find</option>
        <option value="9">ワーシャルフロイド法</option>
        <option value="10">ユークリッドの互除法</option>
        <option value="11">整数問題</option>
        <option value="12">セグ木</option>
        <option value="13">論理演算</option>
        <option value="14">包除原理</option>
        </select>
      </div>
    </div>

      <div className="flex justify-end items-center gap-x-2 py-3 px-4 border-t">
        <button type="button" className="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none" data-hs-overlay="#hs-toggle-password-modal-example">
          Close
        </button>
        <button onClick={handleSaveChanges} className="py-2 px-3 inline-flex items-center gap-x-2 text-sm font-semibold rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none">
      Save changes
    </button>
      </div>
    </div>
  </div>
</div>



<br></br>


<div className="flex flex-col">
    <div className="-m-1.5 overflow-x-auto">
      <div className="p-1.5 min-w-full inline-block align-middle">
        <div className="border rounded-lg divide-y divide-gray-200">

        <div className="overflow-hidden">
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
            <tr>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Name</th>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Difficulty</th>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Category</th>
                  <th scope="col" className="px-6 py-3 text-start text-xs font-medium text-gray-500 uppercase">Link</th>
            </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {problems.map((problem, index) => (
                <tr key={index}>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{problem.problem_id}</td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{problem.difficulties}</td>
                  <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-800">{problem.classifier}</td>
                  <td className="px-6 py-4 whitespace-nowrap text-start text-sm font-medium">
                    <a href={problem.url}   target="_blank" rel="noopener noreferrer" className="text-blue-600 hover:text-blue-800 disabled:opacity-50 disabled:pointer-events-none">Learn</a>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        </div>
      </div>
    </div>
  </div>

  
    </div>
  );
};

export default Searchbar;


{/* <button type="button" className="button-style" onClick={handleOpenModal}>
  Select Algorithm
</button>

<br></br>

{showModal && (
  <div id="hs-toggle-password-modal-example" className="modal-overlay">
      <div className="modal-content">
          <div className="modal-header">
              <h3 className="modal-title">
                  検索したいアルゴリズムを選択してください。
              </h3>
              <button type="button" className="close-button" onClick={handleCloseModal}>
                  <span className="sr-only">Close</span>
                  <svg className="flex-shrink-0 w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
              </button>
          </div>
          <div className="modal-body">
          <option value="Name">Select the option ...</option>
          <option value="1">全探索</option>
          <option value="2">累積和及びimos法</option>
          <option value="3">二分探索</option>
          <option value="4">座標圧縮、ランレングス圧縮</option>
          <option value="5">動的計画法</option>
          <option value="6">尺取り法</option>
          <option value="7">BFS</option>
          <option value="15">DFS</option>
          <option value="8">Union-Find</option>
          <option value="9">ワーシャルフロイド法</option>
          <option value="10">ユークリッドの互除法</option>
          <option value="11">整数問題</option>
          <option value="12">セグ木</option>
          <option value="13">論理演算</option>
          <option value="14">包除原理</option>
          </div>
          <div className="modal-footer">
              <button type="button" className="close-button-style" onClick={handleCloseModal}>
                  Close
              </button>
              <button onClick={handleSaveChanges} className="save-button-style">
                  Save changes
              </button>
          </div>
      </div>
  </div>
)} */}