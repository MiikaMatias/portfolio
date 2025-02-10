import { useState } from 'react'
import './App.css'
import NavBar from './components/NavBar'
import Intro from './components/Introduction'
import Portfolio from './components/Portfolio'

function App() {
  return (  
    <>
      <div className="main-content">
        <Intro></Intro>
        <Portfolio></Portfolio>
      </div>
    </>
  )
}

export default App
