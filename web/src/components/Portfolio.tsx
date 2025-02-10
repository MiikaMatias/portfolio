
import React, {ChangeEvent, useState} from 'react';

function Portfolio() {

    const [query, setQuery] = useState("")
 
    async function fetchReply() {
        
    }

    var input = document.getElementById("query");
    input?.addEventListener('keypress', (event) => {      
        if (event.key === 'Enter' && event.target instanceof HTMLInputElement && event.target.value!="") {
            console.log('Sending a request for:'+event.target.value);
            setQuery("");
        }
    });

    return(
        <h1 className="portfolio">
            <a>Hey, I'm Miika's porfolio! Ask me anything...</a>
            <input id="query" type="text" value={query}
                    onChange={(changeEvent: ChangeEvent<HTMLInputElement>) => {
                        setQuery(changeEvent.target.value)
                    }}></input>
        </h1>        
    )
}

export default Portfolio;