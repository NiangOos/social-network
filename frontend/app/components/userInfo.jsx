export const Userinfo = () => {
    return (
        <div className='rightt w-full h-full absolute top-0 right-0 z-10 hidden transition-transform'>
            <div className='w-[360px] border-l shadow-lg h-screen bg-white absolute right-0 top-0 uk-animation-slide-right-medium delay-200 z-50 dark:bg-dark2 dark:border-slate-700'>
                <div className='w-full h-1.5 bg-gradient-to-r to-purple-500 via-red-500 from-pink-500 -mt-px' />
                <div className='py-10 text-center text-sm pt-20'>
                    <img src='assets/images/avatars/avatar-3.jpg' className='w-24 h-24 rounded-full mx-auto mb-3' />
                    <div className='mt-8'>
                        <div className='md:text-xl text-base font-medium text-black dark:text-white'>
                            {" "}
                            Monroe Parker
                        </div>
                        <div className='text-gray-500 text-sm mt-1 dark:text-white/80'>@Monroepark</div>
                    </div>
                    <div className='mt-5'>
                        <a
                            href='#'
                            className='inline-block rounded-full px-4 py-1.5 text-sm font-semibold bg-secondery'
                        >
                            View profile
                        </a>
                    </div>
                </div>
                <hr className='opacity-80 dark:border-slate-700' />
                <ul className='text-base font-medium p-3'>
                    <li>
                        <div className='flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery'>
                            <ion-icon name='notifications-off-outline' className='text-2xl' /> Mute Notification
                            <label className='switch cursor-pointer ml-auto'>
                                {" "}
                                <input type='checkbox' defaultChecked />
                                <span className='switch-button !relative' />
                            </label>
                        </div>
                    </li>
                    <li>
                        {" "}
                        <button
                            type='button'
                            className='flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery'
                        >
                            {" "}
                            <ion-icon name='flag-outline' className='text-2xl' /> Report{" "}
                        </button>
                    </li>
                    <li>
                        {" "}
                        <button
                            type='button'
                            className='flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery'
                        >
                            {" "}
                            <ion-icon name='settings-outline' className='text-2xl' /> Ignore messages{" "}
                        </button>{" "}
                    </li>
                    <li>
                        {" "}
                        <button
                            type='button'
                            className='flex items-center gap-5 rounded-md p-3 w-full hover:bg-secondery'
                        >
                            {" "}
                            <ion-icon name='stop-circle-outline' className='text-2xl' /> Block{" "}
                        </button>{" "}
                    </li>
                    <li>
                        {" "}
                        <button
                            type='button'
                            className='flex items-center gap-5 rounded-md p-3 w-full hover:bg-red-50 text-red-500'
                        >
                            {" "}
                            <ion-icon name='trash-outline' className='text-2xl' /> Delete Chat{" "}
                        </button>{" "}
                    </li>
                </ul>
                {/* close button */}
                <button
                    type='button'
                    className='absolute top-0 right-0 m-4 p-2 bg-secondery rounded-full'
                    uk-toggle='target: .rightt ; cls: hidden'
                >
                    <ion-icon name='close' className='text-2xl flex' />
                </button>
            </div>
            {/* overly */}
            <div
                className='bg-slate-100/40 backdrop-blur absolute w-full h-full dark:bg-slate-800/40'
                uk-toggle='target: .rightt ; cls: hidden'
            />
        </div>
    );
};
