import { useState } from 'react'
import './App.css'
import NavBar from './components/NavBar'

function App() {
  return (  
    <>
      <NavBar></NavBar>
      <div class="main-content">
        <h1 className="introduction">
          <a>Hey, I'm Miika's porfolio! Ask me anything...</a>
          <input></input>
        </h1>        
      </div>
    </>
  )
}

export default App
