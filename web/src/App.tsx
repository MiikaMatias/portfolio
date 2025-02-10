import { useState } from 'react'
import './App.css'
import NavBar from './components/NavBar'
import Intro from './components/Introduction'
import Portfolio from './components/Portfolio'
import AnswerBox from './components/AnswerBox.tsx';

function App() {
  return (  
    <>
      <div className="main-content">
        <Intro></Intro>
        <Portfolio></Portfolio>
        <AnswerBox></AnswerBox>
      </div>
    </>
  )
}

export default App
