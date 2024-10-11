<script>
  import BellSvg from './BellSvg.svelte'; 
  import LogoutSvg from './LogoutSvg.svelte';
  import Notification from './Notification.svelte';
  import ProfileSvg from './ProfileSvg.svelte';
  import Group from './group/+page.svelte'
//   import { bellIcon } from './utils.js'; // Adjust the path if needed

// import { Dropdown, DropdownItem, DropdownDivider, DropdownHeader } from 'flowbite-svelte';

export let selectedPostComments = [];
export let showNotification = false;
export let showGroupInfo = false;


let isOpen = false;

function toggleDropdown() {
    isOpen = !isOpen;
}


function bellIcon() {
  showNotification = true;
  showGroupInfo = false

  selectedPostComments = [];
  console.log("clickkkkk");
}
    


</script>


<section id="commentSection" class="bg-white flex flex-col min-h-screen w-[35%] w-max-[40vw]">
    <div class="flex flex-row">
        <input class="text-xxl  text-blue-500 m-4  max-w-[px] p-1" placeholder="Search...">
        <div class="flex items-center align-middle">
            <ul class="flex flex-row items-center h-[7vh] justify-center">
                <li class="mr-2 cursor-pointer">
                    <button on:click={toggleDropdown} class="bg-yellow-400 rounded-full z-100 relative p-1 w-[32px] m-1">O</button>
                    {#if isOpen}
                    <ul class="flex flex-col absolute z-10 right-0 mt-2 w-48 bg-[#6C6C6C] rounded shadow-lg">
                        <li class="px-4 py-2 hover:bg-gray-200">
                            <div class="flex items-center">
                                <ProfileSvg/>
                                <a href="/profile" class="text-white">Profile</a>
                            </div>
                        </li>
                        <li class="px-4 py-2 hover:bg-gray-200">
                            <div class="flex items-center">
                                <LogoutSvg/>
                                <a href="/logout" class="text-white">Logout</a>
                            </div>
                        </li>
                    </ul>
                    {/if}
                </li>
                <button on:click={bellIcon} class="ml-2 z-100 cursor-pointer"> 
                    <BellSvg/>
                </button>
            </ul>
        </div>
    </div>
    <div class="h-screen overflow-auto">
        <!-- <button on:click={hideSectionComment}>X</button> -->
        <div class="flex flex-col">
            <!-- Render the comments for the selected post -->
            {#if selectedPostComments.length > 0}
                    <span class="font-bold ml-2">Commentaires</span>

                {#each selectedPostComments as comment}
                    <div class="flex flex-col bg-gray-300 h-auto p-1 m-3 text-xs text-center">
                        <div class="flex flex-col justify-around p-4 w-[90%]">  
                            <div><strong>Name:</strong> {comment.name}</div>
                            <div><strong>Email:</strong> {comment.email}</div>
                            <div><strong>Body:</strong> {comment.body}</div>
                        </div>
                    </div>
                {/each}
            {:else if showNotification === true}
                 <span class="font-bold ml-2">Notifications</span>
                <Notification/>
            {:else if showGroupInfo === true}
                <section class="flex flex-col h-screen">
                    <div class=" flex flex-col h-[15vh] mb-8 relative">
                       <div class="bg-[#0094FF] h-[60%] flex relative"></div>
                            <div class="flex  absolute top-8 left-10">
                                <img class="h-[120px] w-[120px] rounded-lg bg-slate-400" alt="Image for the group">
                                <span class="flex text-sm ml-2 h-[3vh] items-center font-bold relative top-16">Founded by : USERNAME</span>
                            </div>
                       </div>
                    <div class="flex flex-col h-[30vh]">
                        <span class="flex font-bold ml-2">About</span>
                        <div class="flex justify-center">
                            <textarea disabled class="flex text-center resize-none w-96 h-52 m-2 rounded-lg bg-customGray"></textarea>
                        </div>
                    </div>
                    <div class="flex flex-col mt-4 mb-4 w-full ">
                    <span class="font-bold ml-2">Membres</span>

                        <div class="flex flex-col">
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 1</p>
                            </div>
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 2</p>
                            </div>
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 3</p>
                            </div>
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 4</p>
                            </div>
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 5</p>
                            </div>
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 6</p>
                            </div>
                            <div class="flex flex-row m-2">
                                <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                                <p class="bg-customGray flex items-center w-full p-2">Membre 7</p>        
                            </div>
                        </div>
                    </div>
                </section>
            {:else}
            <section class="flex flex-col h-screen">
            <div class=" flex flex-col h-[15vh] mb-8 relative">
               <div class="bg-[#0094FF] h-[60%] flex relative"></div>
                    <div class="flex  absolute top-8 left-10">
                        <img class="h-[120px] w-[120px] rounded-lg bg-slate-400" alt="Image for the group">
                        <span class="flex text-sm ml-2 h-[3vh] items-center font-bold relative top-16">Founded by : USERNAME</span>
                    </div>
               </div>
            <div class="flex flex-col h-[30vh]">
                <span class="flex font-bold ml-2">About</span>
                <div class="flex justify-center">
                    <textarea disabled class="flex text-center resize-none w-96 h-52 m-2 rounded-lg bg-customGray"></textarea>
                </div>
            </div>
            <div class="flex flex-col mt-4 mb-4 w-full ">
            <span class="font-bold ml-2">Membres</span>

                <div class="flex flex-col">
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 1</p>
                    </div>
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 2</p>
                    </div>
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 3</p>
                    </div>
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 4</p>
                    </div>
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 5</p>
                    </div>
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 6</p>
                    </div>
                    <div class="flex flex-row m-2">
                        <span class="h-[35px] w-[35px] bg-customPP rounded-sm mr-4 p-2"></span>
                        <p class="bg-customGray flex items-center w-full p-2">Membre 7</p>        
                    </div>
                </div>
            </div>
        </section>
            {/if}
        </div>
    </div>
</section>
