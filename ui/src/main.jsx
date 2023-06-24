import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import history from "./util/history.js";

// const onRedirectCallback = (appState) => {
//     history.push(
//         appState && appState.returnTo ? appState.returnTo : window.location.pathname
//     );
// };

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
