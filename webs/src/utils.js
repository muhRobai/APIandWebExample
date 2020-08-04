
const currentPagePathIs = name => {
    return (window.nav.location.pathname === name)
}

const currentPagePathContains = name => {
    return (window.nav.location.pathname.indexOf(name) > -1)
}

const navigateTo = (page, state) => {
    return function() { window.nav.push(page, state) }
}

const resetNavigation = () => {
    window.nav.replace('/', {})
}

const moveTo = (page, state) => {
    window.nav.replace(page, state)
}

const goBack = _ => {
    window.nav.goBack()
}


const Utils = {
    currentPagePathIs,
    currentPagePathContains,
    navigateTo,
    goBack,
    moveTo,
    resetNavigation,
}

export default Utils