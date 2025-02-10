import pfp from '../assets/pfp.png';

function NavBar() {
    return (
    <div className="leftnav">
        <img src={pfp}></img>
        <a className="right" href="https://github.com/MiikaMatias " target="_blank">Github</a>
        <a className="right" href="https://www.linkedin.com/in/miika-piiparinen-9185221a7/rr" target="_blank">LinkedIn</a>
    </div> 
    )
}

export default NavBar;