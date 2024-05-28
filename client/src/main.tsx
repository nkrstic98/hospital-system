import React from 'react'
import ReactDOM from 'react-dom/client'
// import './index.css'
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { BrowserRouter } from 'react-router-dom';
import App from './App';

export interface RouteAdapterProps {
    children: (props: { history: any; location: any }) => React.ReactNode;
}

ReactDOM.createRoot(document.getElementById('root')!).render(
  <BrowserRouter>
      <App />
  </BrowserRouter>
)
