import React from 'react';
import { useLocation } from 'react-router-dom';

const SideMenu: React.FC = () => {
  const location = useLocation();
  const isAlgorithmPage = location.pathname === '/algorithm';

  return (
<div>
<body className="bg-gray-50">
<div className="sticky top-0 inset-x-0 z-20 bg-white border-y px-4 sm:px-6 md:px-8 lg:hidden">
    <div className="flex items-center py-4">
    <button type="button" className="text-gray-500 hover:text-gray-600" data-hs-overlay="#application-sidebar-brand" aria-controls="application-sidebar-brand" aria-label="Toggle navigation">
        <span className="sr-only">Navigation</span>
        <svg className="w-5 h-5" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
        <path fill-rule="evenodd" d="M2.5 12a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5zm0-4a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5z"/>
        </svg>
    </button>
    <ol className="ms-3 flex items-center whitespace-nowrap" aria-label="Breadcrumb">
        <li className="flex items-center text-sm text-gray-800">
        Syojincoder
        <svg className="flex-shrink-0 mx-3 overflow-visible h-2.5 w-2.5 text-gray-400" width="16" height="16" viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M5 1L10.6869 7.16086C10.8637 7.35239 10.8637 7.64761 10.6869 7.83914L5 14" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
        </svg>
        </li>
        <li className="text-sm font-semibold text-gray-800 truncate" aria-current="page">
        {isAlgorithmPage ? 'Algorithm' : 'Dashboard'}
        </li>
    </ol>
    </div>
</div>

<div id="application-sidebar-brand" className="hs-overlay hs-overlay-open:translate-x-0 -translate-x-full transition-all duration-300 transform hidden fixed top-0 start-0 bottom-0 z-[60] w-64 bg-blue-700 pt-7 pb-10 overflow-y-auto lg:block lg:translate-x-0 lg:end-auto lg:bottom-0 [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-thumb]:rounded-full [&::-webkit-scrollbar-track]:bg-gray-100 [&::-webkit-scrollbar-thumb]:bg-gray-300">
    <div className="px-6">
    <a className="flex-none text-xl font-semibold text-white" href="/" aria-label="Brand">精進コーダー</a>
    </div>

    <nav className="hs-accordion-group p-6 w-full flex flex-col flex-wrap" data-hs-accordion-always-open>
    <ul className="space-y-1.5">
        <li>
        <a className={`flex items-center gap-x-3 py-2 px-2.5 text-sm text-white rounded-lg ${isAlgorithmPage ? 'hover:bg-blue-600-300' : 'bg-blue-600'}`} href="/">
            <svg className="flex-shrink-0 w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" ><path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><polyline points="9 22 9 12 15 12 15 22"/></svg>
            ダッシュボード
        </a>
        </li>

        <li>
        <a className={`flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-white rounded-lg ${isAlgorithmPage ? 'bg-blue-600' : 'hover:bg-blue-600-300'}`} href="/algorithm">
        <svg className="flex-shrink-0  w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="10" x="3" y="11" rx="2"/><circle cx="12" cy="5" r="2"/><path d="M12 7v4"/><line x1="8" x2="8" y1="16" y2="16"/><line x1="16" x2="16" y1="16" y2="16"/></svg>
        アルゴリズムの検索
        </a></li>
        <li><a className="w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-white hover:text-white rounded-lg hover:bg-blue-600-300" href="#">
        <svg className="flex-shrink-0 w-4 h-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"/><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/></svg>
        ドキュメント (準備中)
        </a></li>
    </ul>
    </nav>
</div>
</body>
</div>
  );
};

export default SideMenu;