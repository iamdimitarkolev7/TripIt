const cookies = document.cookie.split('; ')
                               .reduce((acc, cookie) => {
                                    const [cookieName, cookieValue] = cookie.split('=');
                                    acc[cookieName] = cookieValue;
                                    return acc;}, {});

const isLoggedIn = !!cookies['token'];

export default isLoggedIn;