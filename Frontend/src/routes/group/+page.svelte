<script>
   import { goto } from '$app/navigation'; // Import the SvelteKit navigation function
//   import Notification from '../Notification.svelte';
  import Group from '../Group.svelte'
//   import Comment from '../Comment.svelte'
  import CreatePostSvg from '../CreatePostSvg.svelte';
  import EditSvg from '../EditSvg.svelte';
  import MessageSvg from '../MessageSvg.svelte';
   import Sidebar from '../Sidebar.svelte';
   import { onMount } from 'svelte';
  import PeopleSvg from '../PeopleSvg.svelte';
//    import { bellIcon } from '../utils.js'; // Adjust the path if needed


    let users = [];
    let posts = []
    let comments = []
    let selectedPostComments = []; 
    let albums = [];
    let photos = [];
    let showNotification = false;
    let showGroupInfo = false;


    // Fetch data when the component is mounted
    const fetchUsers = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/users');
        users = await response.json();
    };

    const fetchPosts = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/posts')
        posts = await response.json();
        console.log("posts: ", posts);
    };

    const fetchComments= async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/comments')
        comments = await response.json();
        console.log("comments: ", comments);
    };

    const fetchAlbums = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/albums');
        albums = await response.json();
    };
    const fetchPhotos = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/photos');
        photos = await response.json();
    };


    // function getCookieValue() {
    //     let theCookies = document.cookie.split(";"); // Split cookies by semicolon
    //     for (let i = 0; i < theCookies.length; i++) {
    //         let cookie = theCookies[i].trim(); // Remove leading spaces
    //         console.log("cookie is :",cookie);
    //         if (cookie.startsWith("sessionId=")) {
    //             return cookie.substring("sessionId=".length); // Get the value after "sessionId="
    //         }
    //     }
    //     return null; // Return null if sessionId cookie is not found
    // }

    //   let sessionId = null;

    //   onMount(() => {
    //     sessionId = getCookieValue(); // Only call this when the DOM is ready
    //     console.log("sessionId:", sessionId); // This will now log the sessionId
    //   });


    // Function to safely check if we are in the browser environment
    function isBrowser() {
        return typeof window !== 'undefined' && typeof document !== 'undefined';
    }

    // Function to get the cookie value, only run in the browser
    function getCookieValue() {
      if (isBrowser()) {
        console.log(document.cookie); // Safe access after the page is loaded
        let theCookies = document.cookie.split(";"); // Split cookies by semicolon
        for (let i = 0; i < theCookies.length; i++) {
          let cookie = theCookies[i].trim(); // Remove leading spaces
          console.log("cookie is:", cookie);
          if (cookie.startsWith("sessionId=")) {
            return cookie.substring("sessionId=".length); // Get the value after "sessionId="
          }
        }
      }
      return null; // Return null if sessionId cookie is not found
    }

    
    let cookieValue =  getCookieValue();
    console.log("Value of the cookie: ", cookieValue);


//     async function fetchCookieValue() {

//        try {
//            const response = await fetch('http://localhost:8080/verificationSessionId', {
//                method: 'POST',
//                body: JSON.stringify(cookieValue)
//            });

//            if (response.ok) {
//                const result = await response.json();
//                 if (result.hasOwnProperty("Success")) {
//                 console.log('File and data uploaded successfully:', result);
//                 //Setting the cookie if the user is already registered

//                } else {
//                 goto('/');
//                 return
//                }
//            } else {
//                console.error('Cookie not valid');
//            }
//        } catch (error) {
//            console.error('Error with the cookie:', error);
//        }
//    }

  
  export function deleteCookie() {
      var res = document.cookie;
      var arrayCookie = res.split(";");
      for(var i = 0; i < arrayCookie.length; i++) {
          var key = arrayCookie[i].split("=");
          document.cookie = key[0]+" =; expires = Thu, 01 Jan 1970 00:00:00 UTC";
      }
      goto('/')
   }


    // fetchCookieValue()
    fetchUsers();
    fetchPosts()
    fetchComments()
    fetchAlbums();
    fetchPhotos();


    let previousUserId = null;

    const findUserName = (userId) => {
        const user = users.find(user => user.id === userId);
        return user ? user.name : 'Unknown User';
    };

    const findCommentPosts = (postId) => {
        return comments.filter(comment => comment.postId === postId);
    }

    function onClick(event) {
        event.preventDefault()
        console.log("Hello");
    }

    function displaySectionComment(postId) {
        selectedPostComments = findCommentPosts(postId);
        let sectionComment = document.getElementById("commentSection")
        sectionComment.style.display = "flex"
    }


    function displayGroup() {
        console.log("clickkkkk");
        showGroupInfo = true
        let sectionComment = document.getElementById("commentSection")
        sectionComment.style.display = "flex"
    }


    function hideSectionComment() {
        let sectionComment = document.getElementById("commentSection")
        sectionComment.style.display = "none"
        selectedPostComments = []
    }

    function bellIcon() {
        console.log("clickkkkk");
        selectedPostComments = []; 
        showNotification = true
        showGroupInfo = false

    }

    function peopleIcon() {
        console.log("clickkkkk");
        selectedPostComments = []; 
        showNotification = false
        showGroupInfo = true
        
    }


</script>

<section class="flex min-h-screen bg-gray-200">
    <Sidebar { users } { albums } { photos } />
    <section class="flex min-h-screen w-full">
        <section class="flex flex-col min-h-screen w-[80%]">
            <section>
                <div class="bg-white flex justify-between items-center mt-4 m-2.5 p-0.5">
                    <div class="flex items-center align-middle font-bold text-3xl text-blue-500 t-2 ml-2">
                        <input placeholder="New group">
                    </div>
                    <div class="flex justify-center items-center mr-2"> 
                    <EditSvg/>
                    <CreatePostSvg/>
                    <MessageSvg/>
                    <button on:click={peopleIcon} class="cursor-pointer">
                        <PeopleSvg/>
                    </button>
                    </div>
                </div>
                <!-- <textarea class="h-7 w-96 m-2" placeholder="Ecrire le contenu du post"></textarea> -->
                    <div class="m-12 max-h-[100vh] overflow-auto flex flex-col items-center   
                          [&::-webkit-scrollbar]:w-2
                          [&::-webkit-scrollbar-track]:bg-gray-100
                          [&::-webkit-scrollbar-thumb]:bg-gray-300
                          [&::-webkit-scrollbar-track]:rounded-full
                          [&::-webkit-scrollbar-thumb]:rounded-full">
                        <!-- Render the fetched posts -->
                        <span>Create the first post</span>
                        {#if posts.length > 0}
                            {#each posts as post}
                                <button on:click={displaySectionComment(post.id)} id="postDiv" class="flex flex-col bg-white h-auto p-2 mt-4 mb-4 hover:scale-103 w-[95%]">
                                    <div class="flex flex-row justify-around p-4">
                                        <div>Username: {findUserName(post.userId)}</div>
                                        <div class="flex">
                                            <div class="w-4 h-5 bg-red-500"></div>
                                            <div class="w-4 h-5 bg-red-900"></div>
                                        </div>
                                        <div class="flex">
                                            <div class="w-4 h-5 bg-red-500"></div>
                                            <div class="w-4 h-5 bg-red-900"></div>
                                        </div>
                                        <!-- <div>Post ID: {post.id}</div> -->
                                    </div>
                                    <div class="flex p-2">Title: {post.title}</div>
                                    <div class="flex bg-gray-300 p-4">Body: {post.body}</div>
                                </button>
                            {/each}
                        {:else} Comment
                        <p>Loading...</p>
                        {/if}
                    </div>
            </section>
        </section>
            <Group {bellIcon} {selectedPostComments} {showNotification} {peopleIcon} {showGroupInfo}  /> <!--Put the "showGroupInfo" oin the Group tag to use it in the Group.svelte file-->
    </section>
</section>


