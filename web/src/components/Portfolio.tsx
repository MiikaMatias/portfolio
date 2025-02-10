
import React, {ChangeEvent, useState} from 'react';
import type { paths, components } from "../api/api.d.ts";


type SuccessResponse =
  paths["/rag-reply"]["get"]["responses"][200]["content"]["application/json"];


function Portfolio() {

    const [query, setQuery] = useState("");
 
    async function fetchReply(queryString: string) {
        const requestResponse = await fetch(`http://localhost:8889/rag-reply?prompt=${queryString}`, {
            method: "GET",
            headers: { "Content-Type": "application/json" },
        });

        if (!requestResponse.ok) {
            throw new Error("Failed to fetch data");
        }

        const data: SuccessResponse = await requestResponse.json();
        console.log(data);

        var answer = document.getElementById("answer-new");
        var prevQuery = document.getElementById("answer-new");
        if (answer === null) {
            throw new Error("Could not find answer element, which is weird")
        }
        var prevQuery = document.getElementById("prevQuery");
        if (prevQuery === null) {
            throw new Error("Could not find prevQuery, which is weird")
        }

        answer.textContent = data.RagReply;
        prevQuery.textContent = queryString;
    }

    var input = document.getElementById("query");
    input?.addEventListener('keypress', (event) => {      
        if (event.key === 'Enter' && event.target instanceof HTMLInputElement && event.target.value!="") {
            fetchReply(event.target.value);
            console.log('Sending a request for:'+event.target.value);
            setQuery("");
        }
    });

    return(
        <h1 className="portfolio">
            <a id="prevQuery">Hey, I'm Miika's portfolio! Ask me anything...</a>
            <input id="query" type="text" value={query}
                    onChange={(changeEvent: ChangeEvent<HTMLInputElement>) => {
                        setQuery(changeEvent.target.value)
                    }}>
            </input>
        </h1>        
    )
}

export default Portfolio;