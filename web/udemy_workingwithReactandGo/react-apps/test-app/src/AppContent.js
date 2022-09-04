import React, { Component } from 'react';

export default class AppContent extends Component {

    fetchList = () => {
        fetch('https://jsonplaceholder.typicode.com/posts')
        .then((response) => response.json())
        .then(json => {
            console.log(json)
            let posts = document.getElementById("post-list");
            json.forEach(function(obj){
                let li = document.createElement("li");
                li.appendChild(document.createTextNode(obj.title))
                posts.appendChild(li)
            }
            );
        })

    }


    render(){
        return (
            <p>
                This is my content!
                <br/>
                <button onClick={this.fetchList} class="btn btn-primary">Fetch</button>
            
                <hr />

                <ul id="post list"></ul>
            </p>
        );
    }
}
