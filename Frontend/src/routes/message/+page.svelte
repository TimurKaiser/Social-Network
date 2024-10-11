<script>
   import Sidebar from '../Sidebar.svelte';

    let users = [];
    let posts = []
    let comments = []
    let selectedPostComments = []; 
    let albums = [];
    let photos = [];


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

    fetchUsers();
    fetchPosts()
    fetchComments()
    fetchAlbums();
    fetchPhotos();

</script>


<section class="flex min-h-screen bg-gray-100">
    <Sidebar { users } { albums } { photos } />
    <section class="flex min-h-screen w-full">
        <section class="flex flex-col min-h-screen w-full">
            <div class="flex flex-row justify-between m-3 p-2 h-[8vh]">
                <div class="flex items-center align-middle font-bold text-3xl text-blue-500">UserName</div>
                <div class="flex items-center align-middle">
                    <ul class="flex flex-row items-center w-[12vw] h-[7vh] justify-center">
                        <li class="mr-2">
                            <span class="bg-yellow-400 rounded-full p-1 w-[32px] m-1">O</span>
                        </li>
                        <li class="ml-2"> 
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M4 19V17H6V10C6 8.61667 6.41667 7.39167 7.25 6.325C8.08333 5.24167 9.16667 4.53333 10.5 4.2V3.5C10.5 3.08333 10.6417 2.73333 10.925 2.45C11.225 2.15 11.5833 2 12 2C12.4167 2 12.7667 2.15 13.05 2.45C13.35 2.73333 13.5 3.08333 13.5 3.5V4.2C14.8333 4.53333 15.9167 5.24167 16.75 6.325C17.5833 7.39167 18 8.61667 18 10V17H20V19H4ZM12 22C11.45 22 10.975 21.8083 10.575 21.425C10.1917 21.025 10 20.55 10 20H14C14 20.55 13.8 21.025 13.4 21.425C13.0167 21.8083 12.55 22 12 22ZM8 17H16V10C16 8.9 15.6083 7.95833 14.825 7.175C14.0417 6.39167 13.1 6 12 6C10.9 6 9.95833 6.39167 9.175 7.175C8.39167 7.95833 8 8.9 8 10V17Z" fill="#1D1B20"/>
                            </svg>
                        </li>
                    </ul>
                </div>
            </div>
            <section class="flex flex-col">
                <div class="flex flex-col w-auto m-2.5 h-screen overflow-auto">
                    <div class="bg-customGray flex justify-center self-start h-40 m-5 p-4 items-center w-[30vw]">
                        <sapn class="bg-gray-300 float-left">Message</sapn>
                    </div>
                    <div class="bg-customGray flex justify-center self-end h-24 m-5 p-4 items-center w-[30vw]">
                        <sapn class="float-start">Message</sapn>
                    </div>
                    <div class="bg-customGray flex justify-center self-start h-16 m-5 p-4 items-center w-[30vw]">
                        <sapn class="float-end">Message</sapn>
                    </div>
                    <div class="bg-customGray flex justify-center self-start h-16 m-5 p-4 items-center w-[30vw]">
                        <sapn class="float-start">Message</sapn>
                    </div>
                    <div class="bg-customGray flex justify-center self-end h-16 m-5 p-4 items-center w-[30vw]">
                        <sapn class="float-end">Message</sapn>
                    </div>        
                </div>
                <div class="flex justify-center items-center">
                    <textarea class="bg-customGray flex text-center h-12 text-black placeholder-black w-[90%] resize-none" placeholder="Ecrivez votre message..."></textarea>
                    <button class="bg-customGray ml-4 p-2">Send</button>
                </div>
            </section>
        </section>
    </section>
</section>


