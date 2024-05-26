<template>
    <header>
        <nav class="navbar bg-bg shadow-md">
            <div class="navbar-start">
                <div class="dropdown">
                    <button tabindex="0" role="button" class="btn btn-ghost sm:hidden">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h8m-8 6h16" /></svg>
                    </button>
                    <ul tabindex="0" class="text-link menu menu-sm dropdown-content mt-3 z-[1] p-2 bg-secondary rounded-box w-52">
                    <li><a href="/about">About</a></li>
                    <li>
                        <p class="text-neutral">Login or Signup</p>
                        <ul class="p-2">
                            <li><a href="/login">Login</a></li>
                            <li><a href="/signup">Signup</a></li>
                        </ul>
                    </li>
                </ul>
                </div>
                <a class="btn btn-ghost text-xl text-accent" href="/">Playground</a>
            </div>
            <div class="navbar-end hidden sm:flex">
                <div class="menu menu-horizontal text-lg text-link items-center space-x-2">
                    <li><a href="/about" class="hover:text-text">About</a></li>
                    <li><a href="/login" class="btn btn-accent">Login</a></li>
                </div>
            </div>
        </nav>
    </header>
</template>