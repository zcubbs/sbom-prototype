// import LogoDark from '../assets/logo_black.svg';
// import LogoWhite from '../assets/logo_white.svg';
// @ts-ignore
import LogoDefault from '../assets/logo.png';

export function Logo({colorScheme}) {

    const css = {
        height: '32px',
    }

    return (
        <>
            {colorScheme === 'dark' ?
                <img style={css} src={LogoDefault} alt="Logo"/>
                :
                <img style={css} src={LogoDefault} alt="Logo"/>
            }
        </>
    );
}
