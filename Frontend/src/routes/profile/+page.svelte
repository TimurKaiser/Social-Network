<script>
    import { goto } from '$app/navigation'; // Import the SvelteKit navigation function
    import Comment from '../Comment.svelte'
    import RightPanelFunctions from '../RightPanelFunctions.svelte'
    import Profile from '../Profile.svelte'
    import Sidebar from '../Sidebar.svelte';
    import { browser } from '$app/environment'; // Import to check for browser environment
    import Notification from '../Notification.svelte';

    let users = [];
    let posts = [];
    let comments = [];
    let albums = [];
    let photos = [];
    let cookieValue = null;
    let selectedPostComments = [];
    let showNotification = false;
    let showProfile = true

    // Fetch data when the component is mounted
    const fetchUsers = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/users');
        users = await response.json();
    };

    const fetchPosts = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/posts');
        posts = await response.json();
        // console.log("posts: ", posts);
    };

    const fetchComments = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/comments');
        comments = await response.json();
        // console.log("comments: ", comments);
    };

    const fetchAlbums = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/albums');
        albums = await response.json();
    };

    const fetchPhotos = async () => {
        const response = await fetch('https://jsonplaceholder.typicode.com/photos');
        photos = await response.json();
    };

    function getCookieValue() {
        if (browser) {
            // Ensure the document is only accessed in the browser
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

    cookieValue = getCookieValue();
    console.log("Value of the cookie:", cookieValue);

    async function fetchCookieValue() {
        if (!cookieValue) {
            return;
        }
        try {
            const response = await fetch('http://localhost:8080/verificationSessionId', {
                method: 'POST',
                body: JSON.stringify(cookieValue)
            });

            if (response.ok) {
                const result = await response.json();
                if (result.hasOwnProperty("Success")) {
                    console.log('File and data uploaded successfully:', result);
                } else {
                    goto('/');
                }
            } else {
                console.error('Cookie not valid');
            }
        } catch (error) {
            console.error('Error with the cookie:', error);
        }
    }

    // Fetch data after checking for browser environment
    if (browser) {
        fetchCookieValue();
        fetchUsers();
        fetchPosts();
        fetchComments();
        fetchAlbums();
        fetchPhotos();
    }

    let previousUserId = null;

    const findUserName = (userId) => {
        const user = users.find(user => user.id === userId);
        return user ? user.name : 'Unknown User';
    };

    const findCommentPosts = (postId) => {
        return comments.filter(comment => comment.postId === postId);
    };

    function onClick(event) {
        event.preventDefault();
        console.log("Hello");
    }

    function displaySectionComment(postId) {
        selectedPostComments = findCommentPosts(postId);
        showNotification = false
        showProfile = false
        if (browser) {
            let sectionComment = document.getElementById("commentSection");
            sectionComment.style.display = "flex";
        }
    }

    function hideSectionComment() {
        if (browser) {
            let sectionComment = document.getElementById("commentSection");
            sectionComment.style.display = "none";
            selectedPostComments = [];
        }
    }

    let isEditing = false;

    function toggleEditMode() {
    isEditing = !isEditing;
  }

  function saveChanges() {
    // Simulate saving the data to an API
    fetch(`https://jsonplaceholder.typicode.com/users/1`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(users[0])
    }).then(() => {
      isEditing = false;
      originalUser = { ...users[0] }; // Update originalUser after saving
    });
  }

  function cancelEdit() {
    user = { ...originalUser }; // Revert to original data
    isEditing = false;
  }

  function handleToggleProfile(event) {
        const { showProfile: updatedShowProfile, showNotification: updatedShowNotification, selectedPostComments: updatedSelectedPostComments } = event.detail;
        showProfile = updatedShowProfile;
        showNotification = updatedShowNotification;
        selectedPostComments = updatedSelectedPostComments
    }
</script>
 
 
 <section class="flex min-h-screen bg-gray-100">
    <Sidebar { users } { albums } { photos } />
     <section class="flex min-h-screen w-full">
         <section class="flex flex-col bg-[var(--Light-background,#F8F9FA)] min-h-screen">
            <div class="flex justify-between items-center flex-shrink-0">
                <span class="ms-[40px] ml-2 text-[#007BFF] text-2xl font-inter font-bold break-words m-4 max-w-[px] p-2">
                    {#each users as user}
                    {#if user.id == 1}
                    {user.name}
                    {/if}
                    {/each}
                </span>
                <button>
                <svg xmlns="http://www.w3.org/2000/svg" width="32" height="33" viewBox="0 0 32 33" fill="none">
                    <path d="M1.33337 31.1666L3.93337 22.2333C3.51115 21.3222 3.19449 20.3888 2.98337 19.4333C2.77226 18.4777 2.66671 17.5 2.66671 16.5C2.66671 14.6555 3.01671 12.9222 3.71671 11.3C4.41671 9.67774 5.36671 8.26663 6.56671 7.06663C7.76671 5.86663 9.17782 4.91663 10.8 4.21663C12.4223 3.51663 14.1556 3.16663 16 3.16663C17.8445 3.16663 19.5778 3.51663 21.2 4.21663C22.8223 4.91663 24.2334 5.86663 25.4334 7.06663C26.6334 8.26663 27.5834 9.67774 28.2834 11.3C28.9834 12.9222 29.3334 14.6555 29.3334 16.5C29.3334 18.3444 28.9834 20.0777 28.2834 21.7C27.5834 23.3222 26.6334 24.7333 25.4334 25.9333C24.2334 27.1333 22.8223 28.0833 21.2 28.7833C19.5778 29.4833 17.8445 29.8333 16 29.8333C15 29.8333 14.0223 29.7277 13.0667 29.5166C12.1112 29.3055 11.1778 28.9888 10.2667 28.5666L1.33337 31.1666ZM5.26671 27.2333L9.53337 25.9666C9.84449 25.8777 10.1612 25.8444 10.4834 25.8666C10.8056 25.8888 11.1112 25.9666 11.4 26.1C12.1112 26.4555 12.8556 26.7222 13.6334 26.9C14.4112 27.0777 15.2 27.1666 16 27.1666C18.9778 27.1666 21.5 26.1333 23.5667 24.0666C25.6334 22 26.6667 19.4777 26.6667 16.5C26.6667 13.5222 25.6334 11 23.5667 8.93329C21.5 6.86663 18.9778 5.83329 16 5.83329C13.0223 5.83329 10.5 6.86663 8.43337 8.93329C6.36671 11 5.33337 13.5222 5.33337 16.5C5.33337 17.3 5.42226 18.0888 5.60004 18.8666C5.77782 19.6444 6.04449 20.3888 6.40004 21.1C6.5556 21.3888 6.63893 21.6944 6.65004 22.0166C6.66115 22.3389 6.62226 22.6555 6.53337 22.9666L5.26671 27.2333ZM14.6667 21.8333H17.3334V17.8333H21.3334V15.1666H17.3334V11.1666H14.6667V15.1666H10.6667V17.8333H14.6667V21.8333Z" fill="#6C6C6C"/>
                    </svg>
                </button>
              </div>
                 <section>
                 <!-- <textarea class="ml-[40px] mr-[480px] h-[128px] h-7 w-96 m-2" placeholder="Ecrire le contenu du post"></textarea> -->
                     <div class="m-2.5 h-screen overflow-auto">
                         <!-- Render the fetched posts -->
                         {#if posts.length > 0}
                             {#each posts as post}
                                 <button on:click={displaySectionComment(post.id)} id="postDiv" class="flex flex-col bg-white h-auto p-2 mt-4 mb-4 hover:scale-103 w-full">
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
                         {:else}
                         <p>Loading...</p>
                         {/if}
                     </div>
                 </section>
         </section>
         <section id="commentSection" class="bg-white flex flex-col min-h-screen w-[35%] w-max-[40vw]">
            <RightPanelFunctions {showProfile} {showNotification} {selectedPostComments} on:toggleProfile={handleToggleProfile} />
                {#if showProfile == true && selectedPostComments.length == 0 }
                <Profile {users} />
            {:else if showNotification == true && selectedPostComments.length == 0 }
                <Notification />
            {:else if showProfile == false && showNotification == false}
                <Comment {selectedPostComments} {showNotification} />
            {/if}
        </section>
     </section>
 </section>