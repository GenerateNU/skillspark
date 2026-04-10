import React, { useState } from "react";
import { View, TouchableOpacity } from "react-native";
import { SvgXml } from "react-native-svg";
import { ThemedText } from "@/components/themed-text";
import { Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme";

const CIRCLE_SIZE = 62;
const RING_WIDTH = 2.5;
const RING_GAP = 2; // padding between ring and circle

export const AVATAR_COLORS = [
  "#9CAF7D",
  "#F5D878",
  "#F5A85A",
  "#B5A8E0",
  "#F0A0BC",
  "#E07070",
  "#6B9FD4",
  "#A8D4E8",
  "#B0B0B0",
];

export const DEFAULT_AVATAR_COLOR = AVATAR_COLORS[3]; // lavender purple

const ALIEN_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="60" height="60" rx="30" fill="BGCOLOR"/>
<path d="M19 45C19.5775 45 23.47 48 25.5 48C27.53 48 28.8487 45.2463 31.5 46.5C34.1513 47.7537 34.715 50.0147 40 45" stroke="black" stroke-width="2" stroke-linecap="round"/>
<ellipse cx="43.0595" cy="27.4317" rx="9.33479" ry="13.0687" transform="rotate(40 43.0595 27.4317)" fill="black"/>
<ellipse cx="17.3508" cy="27.657" rx="9.18722" ry="12.8621" transform="rotate(-40 17.3508 27.657)" fill="black"/>
<g filter="url(#filter0_f_5612_6599)">
<path d="M46.0002 22.5C48.0002 24.5 47.5002 27.3333 47.0002 28.5C46.6003 29.7 48.5001 29 49.5 28.5C50 28.1667 51.1 27.3 51.5 26.5C52 25.5 52 22 51 20.5C50.69 20.0349 49.5 19 47.5002 19C46.0131 19 44.0002 20.5 46.0002 22.5Z" fill="white"/>
</g>
<g filter="url(#filter1_f_5612_6599)">
<path d="M16.1685 23.5238C19.0069 24.8413 20.7748 28.215 21.304 29.7372C21.9257 31.2512 22.6895 29.6144 22.9937 28.6068C23.0838 28.0104 23.1832 26.5501 22.8589 25.4805C22.4536 24.1435 19.8521 20.2552 18.0614 19.041C17.5061 18.6644 15.9326 18.0529 14.581 18.9572C13.5759 19.6296 13.3302 22.2063 16.1685 23.5238Z" fill="white"/>
</g>
<defs>
<filter id="filter0_f_5612_6599" x="43.1611" y="17" width="10.6675" height="14.1626" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
<feFlood flood-opacity="0" result="BackgroundImageFix"/>
<feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape"/>
<feGaussianBlur stdDeviation="1" result="effect1_foregroundBlur_5612_6599"/>
</filter>
<filter id="filter1_f_5612_6599" x="11.8877" y="16.4868" width="13.1943" height="15.8647" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
<feFlood flood-opacity="0" result="BackgroundImageFix"/>
<feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape"/>
<feGaussianBlur stdDeviation="1" result="effect1_foregroundBlur_5612_6599"/>
</filter>
</defs>
</svg>`;

const CUTE_FACE_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="60" height="60" rx="30" fill="BGCOLOR"/>
<ellipse cx="8.63883" cy="30.7044" rx="6.66667" ry="4.88263" fill="#D67EA0" fill-opacity="0.5"/>
<ellipse cx="51.5495" cy="30.7044" rx="6.66667" ry="4.88263" fill="#D67EA0" fill-opacity="0.5"/>
<path d="M21.5494 22.2886C21.5494 23.3387 20.6981 24.19 19.648 24.19C18.5979 24.19 17.7466 23.3387 17.7466 22.2886C17.7466 21.2385 18.5979 20.3872 19.648 20.3872C20.6981 20.3872 21.5494 21.2385 21.5494 22.2886Z" fill="black"/>
<path d="M42.2536 22.2886C42.2536 23.3387 41.4023 24.19 40.3522 24.19C39.3021 24.19 38.4508 23.3387 38.4508 22.2886C38.4508 21.2385 39.3021 20.3872 40.3522 20.3872C41.4023 20.3872 42.2536 21.2385 42.2536 22.2886Z" fill="black"/>
<path d="M31.0931 33.4359C31.0931 32.6926 30.7117 32.3209 29.9487 32.3209C29.6113 32.3209 29.2029 32.426 28.7237 32.6363C28.2493 32.8417 27.9754 32.9444 27.9021 32.9444C27.6087 32.9444 27.4619 32.8124 27.4619 32.5483C27.4619 32.1424 27.7554 31.7976 28.3422 31.514C28.9291 31.2303 29.5991 31.0885 30.3522 31.0885C31.8878 31.0885 32.6581 31.6851 32.6629 32.8784C32.6629 33.568 32.4306 34.2257 31.9661 34.8517C31.9465 34.8762 31.856 34.9862 31.6946 35.1818C31.5381 35.3774 31.4599 35.4923 31.4599 35.5266C31.4599 35.6195 31.6066 35.7638 31.9 35.9594C32.1984 36.155 32.4649 36.4289 32.6996 36.781C32.9344 37.1331 33.0517 37.5635 33.0517 38.0721C33.0517 39.0599 32.6898 39.8375 31.9661 40.4048C31.2423 40.9721 30.3791 41.2557 29.3766 41.2557C28.7212 41.2557 28.1442 41.1213 27.6453 40.8523C27.1465 40.5833 26.8971 40.2777 26.8971 39.9353C26.8971 39.593 27.0585 39.4218 27.3813 39.4218C27.5524 39.4218 27.8654 39.5245 28.3202 39.7299C28.775 39.9304 29.2274 40.0307 29.6773 40.0307C30.1321 40.0307 30.5405 39.8693 30.9024 39.5465C31.2643 39.2189 31.4452 38.7934 31.4452 38.2701C31.4452 37.3165 30.9097 36.7125 29.8387 36.4582C29.5306 36.3849 29.3448 36.3335 29.2812 36.3042C29.0709 36.2064 28.9658 36.0425 28.9658 35.8127C28.9658 35.6611 29.1736 35.4826 29.5893 35.2772C30.0099 35.0718 30.3278 34.8517 30.5429 34.617C30.9097 34.2111 31.0931 33.8174 31.0931 33.4359Z" fill="black"/>
</svg>`;

const GAP_TOOTH_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="60.0016" height="60.0016" rx="30.0008" fill="BGCOLOR"/>
<path d="M19.977 24.81C19.977 26.3099 18.7611 27.5258 17.2612 27.5258C15.7613 27.5258 14.5454 26.3099 14.5454 24.81C14.5454 23.3101 15.7613 22.0942 17.2612 22.0942C18.7611 22.0942 19.977 23.3101 19.977 24.81Z" fill="black"/>
<path d="M45.4559 24.81C45.4559 26.3099 44.24 27.5258 42.7401 27.5258C41.2402 27.5258 40.0243 26.3099 40.0243 24.81C40.0243 23.3101 41.2402 22.0942 42.7401 22.0942C44.24 22.0942 45.4559 23.3101 45.4559 24.81Z" fill="black"/>
<path d="M18.3766 35.9182C18.3766 35.4569 18.7505 35.0829 19.2119 35.0829H40.9298C41.3911 35.0829 41.7651 35.4569 41.7651 35.9182C41.7651 36.3796 41.3911 36.7536 40.9298 36.7536H19.2119C18.7505 36.7536 18.3766 36.3796 18.3766 35.9182Z" fill="black"/>
<path d="M12.1162 34.7498H47.1719C47.3742 34.7498 47.538 34.9137 47.5381 35.116C47.5381 40.906 42.8447 45.6003 37.0547 45.6003H22.2344C16.4443 45.6003 11.75 40.9061 11.75 35.116C11.7501 34.9138 11.914 34.7499 12.1162 34.7498Z" fill="white" stroke="black" stroke-width="1.5"/>
<rect x="35" y="35" width="5.22513" height="4.98763" fill="black"/>
</svg>`;

const GLASSES_FACE_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="60" height="60" rx="30" fill="BGCOLOR"/>
<path d="M21.5494 22.2891C21.5494 23.3392 20.6981 24.1905 19.648 24.1905C18.5979 24.1905 17.7466 23.3392 17.7466 22.2891C17.7466 21.239 18.5979 20.3877 19.648 20.3877C20.6981 20.3877 21.5494 21.239 21.5494 22.2891Z" fill="black"/>
<path d="M42.2536 22.2891C42.2536 23.3392 41.4023 24.1905 40.3522 24.1905C39.3021 24.1905 38.4508 23.3392 38.4508 22.2891C38.4508 21.239 39.3021 20.3877 40.3522 20.3877C41.4023 20.3877 42.2536 21.239 42.2536 22.2891Z" fill="black"/>
<path d="M37.3709 30.9511C38.084 30.9512 38.6718 31.5324 38.5659 32.2377C38.2924 34.0589 37.4432 35.7581 36.1252 37.0762C34.5008 38.7006 32.2974 39.613 30.0001 39.613C27.7028 39.613 25.4994 38.7006 23.875 37.0762C22.5569 35.7581 21.7078 34.0589 21.4343 32.2377C21.3283 31.5324 21.9165 30.9511 22.6297 30.9511C23.3427 30.9512 23.907 31.5356 24.0574 32.2327C24.302 33.3668 24.8683 34.4168 25.7013 35.2499C26.8414 36.39 28.3878 37.0304 30.0001 37.0304C31.6124 37.0304 33.1588 36.39 34.2989 35.2499C35.1319 34.4168 35.6982 33.3668 35.9428 32.2327C36.0932 31.5356 36.6578 30.9511 37.3709 30.9511Z" fill="#231717"/>
<g filter="url(#filter0_d_5449_5048)">
<path d="M1.97217 16.3707L6.89847 19.812M6.89847 19.812C6.89847 25.4126 11.4387 29.9528 17.0393 29.9528C22.64 29.9528 27.1802 25.4126 27.1802 19.812M6.89847 19.812C6.89847 14.2114 11.4387 9.67114 17.0393 9.67114C22.64 9.67114 27.1802 14.2114 27.1802 19.812M27.1802 19.812H33.0025M33.0025 19.812C33.0025 25.4126 37.5427 29.9528 43.1434 29.9528C48.744 29.9528 53.2842 25.4126 53.2842 19.812M33.0025 19.812C33.0025 14.2114 37.5427 9.67114 43.1434 9.67114C48.744 9.67114 53.2842 14.2114 53.2842 19.812M53.2842 19.812L58.2106 16.3378" stroke="url(#paint0_linear_5449_5048)" stroke-width="1.50235" shape-rendering="crispEdges"/>
</g>
<defs>
<filter id="filter0_d_5449_5048" x="1.16641" y="8.91992" width="57.8527" height="22.5354" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
<feFlood flood-opacity="0" result="BackgroundImageFix"/>
<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
<feOffset dy="0.375587"/>
<feGaussianBlur stdDeviation="0.187793"/>
<feComposite in2="hardAlpha" operator="out"/>
<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_5449_5048"/>
<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_5449_5048" result="shape"/>
</filter>
<linearGradient id="paint0_linear_5449_5048" x1="30.0914" y1="9.67114" x2="30.0914" y2="29.9528" gradientUnits="userSpaceOnUse">
<stop/>
<stop offset="1" stop-color="#484848"/>
</linearGradient>
</defs>
</svg>`;

const SILLY_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="60.0016" height="60.0016" rx="30.0008" fill="BGCOLOR"/>
<path d="M19.977 24.81C19.977 26.3099 18.7611 27.5258 17.2612 27.5258C15.7613 27.5258 14.5454 26.3099 14.5454 24.81C14.5454 23.3101 15.7613 22.0942 17.2612 22.0942C18.7611 22.0942 19.977 23.3101 19.977 24.81Z" fill="black"/>
<path d="M45.4559 24.81C45.4559 26.3099 44.24 27.5258 42.7401 27.5258C41.2402 27.5258 40.0243 26.3099 40.0243 24.81C40.0243 23.3101 41.2402 22.0942 42.7401 22.0942C44.24 22.0942 45.4559 23.3101 45.4559 24.81Z" fill="black"/>
<path d="M40.0945 35.9182C40.0945 36.9055 39.9649 37.8831 39.713 38.7952C39.4611 39.7073 39.092 40.536 38.6266 41.2341C38.1612 41.9322 37.6087 42.4859 37.0006 42.8637C36.3926 43.2415 35.7408 43.436 35.0827 43.436C34.4245 43.436 33.7728 43.2415 33.1647 42.8637C32.5567 42.4859 32.0042 41.9322 31.5388 41.2341C31.0734 40.536 30.7042 39.7073 30.4523 38.7952C30.2005 37.8831 30.0708 36.9055 30.0708 35.9182H40.0945Z" fill="#E2759F"/>
<path d="M18.3766 35.9182C18.3766 35.4569 18.7505 35.0829 19.2119 35.0829H40.9298C41.3911 35.0829 41.7651 35.4569 41.7651 35.9182C41.7651 36.3796 41.3911 36.7536 40.9298 36.7536H19.2119C18.7505 36.7536 18.3766 36.3796 18.3766 35.9182Z" fill="black"/>
</svg>`;

const SLEEPY_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="60.0016" height="60.0016" rx="30.0008" fill="BGCOLOR"/>
<path d="M23.3936 40H37.3867C37.7252 40 38 40.2748 38 40.6133C37.9998 40.8268 37.8268 40.9998 37.6133 41H23.6064C23.2717 41 23 40.7283 23 40.3936C23.0002 40.1763 23.1763 40.0002 23.3936 40ZM15.9521 27C15.9777 27.0455 15.993 27.0965 15.9971 27.1514L15.6338 27.1758L15.2695 27.1514C15.2736 27.0965 15.2899 27.0455 15.3154 27H15.9521ZM44.6299 27C44.6555 27.0455 44.6707 27.0965 44.6748 27.1514L44.3115 27.1758L43.9473 27.1514C43.9513 27.0965 43.9676 27.0455 43.9932 27H44.6299Z" fill="black" stroke="black" stroke-width="6"/>
<ellipse cx="17.5" cy="41.5" rx="8.5" ry="7.5" fill="url(#paint0_radial_5606_6540)" stroke="#99C0EE"/>
<path d="M48 12H52.6875L48 18H53M43 17H45.8125L43 21H46M39 20H40.875L39 23H41" stroke="black" stroke-width="0.5"/>
<defs>
<radialGradient id="paint0_radial_5606_6540" cx="0" cy="0" r="1" gradientUnits="userSpaceOnUse" gradientTransform="translate(17.5 41.5) rotate(90) scale(7.5 8.5)">
<stop stop-color="#99C0EE" stop-opacity="0.2"/>
<stop offset="1" stop-color="#99C0EE"/>
</radialGradient>
</defs>
</svg>`;

const STAR_FACE_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<g clip-path="url(#clip0_5551_16237)">
<rect width="60" height="60" rx="30" fill="BGCOLOR"/>
<g filter="url(#filter0_d_5551_16237)">
<path d="M13.0437 15.0679C13.5619 13.4732 15.8179 13.4732 16.3361 15.0679L17.7109 19.299C17.9426 20.0122 18.6072 20.4951 19.3571 20.4951H23.8059C25.4827 20.4951 26.1799 22.6407 24.8233 23.6263L21.2241 26.2413C20.6175 26.682 20.3636 27.4633 20.5953 28.1765L21.9701 32.4076C22.4883 34.0023 20.663 35.3284 19.3065 34.3428L15.7073 31.7278C15.1006 31.2871 14.2792 31.2871 13.6725 31.7278L10.0733 34.3428C8.71675 35.3284 6.89155 34.0023 7.4097 32.4076L8.78448 28.1765C9.0162 27.4633 8.76235 26.682 8.15569 26.2413L4.55647 23.6263C3.19994 22.6407 3.8971 20.4951 5.57387 20.4951H10.0227C10.7726 20.4951 11.4372 20.0122 11.6689 19.299L13.0437 15.0679Z" fill="#FFEEC1"/>
<path d="M13.519 15.2222C13.8876 14.0883 15.4923 14.0883 15.8608 15.2222L17.2349 19.4536C17.5335 20.3728 18.3905 20.9946 19.3569 20.9946H23.8062C24.9984 20.9948 25.4943 22.5213 24.5298 23.2222L20.9302 25.8364C20.1484 26.4044 19.8212 27.4115 20.1196 28.3306L21.4946 32.562C21.8631 33.6961 20.5648 34.6389 19.6001 33.938L16.0015 31.3237C15.2196 30.7556 14.1603 30.7556 13.3784 31.3237L9.77979 33.938C8.81511 34.6389 7.51678 33.6961 7.88525 32.562L9.26025 28.3306C9.55865 27.4115 9.23148 26.4044 8.44971 25.8364L4.8501 23.2222C3.88546 22.5213 4.38142 20.9947 5.57373 20.9946H10.0229C10.9893 20.9945 11.8453 20.3726 12.144 19.4536L13.519 15.2222Z" stroke="#C5A576"/>
</g>
<g filter="url(#filter1_d_5551_16237)">
<path d="M43.4168 15.0679C43.9349 13.4732 46.191 13.4732 46.7091 15.0679L48.0839 19.299C48.3156 20.0122 48.9802 20.4951 49.7301 20.4951H54.179C55.8557 20.4951 56.5529 22.6407 55.1964 23.6263L51.5972 26.2413C50.9905 26.682 50.7366 27.4633 50.9684 28.1765L52.3432 32.4076C52.8613 34.0023 51.0361 35.3284 49.6796 34.3428L46.0803 31.7278C45.4737 31.2871 44.6522 31.2871 44.0455 31.7278L40.4463 34.3428C39.0898 35.3284 37.2646 34.0023 37.7827 32.4076L39.1575 28.1765C39.3892 27.4633 39.1354 26.682 38.5287 26.2413L34.9295 23.6263C33.573 22.6407 34.2701 20.4951 35.9469 20.4951H40.3958C41.1457 20.4951 41.8103 20.0122 42.042 19.299L43.4168 15.0679Z" fill="#FFEEC1"/>
<path d="M43.8921 15.2222C44.2606 14.0883 45.8653 14.0883 46.2339 15.2222L47.6079 19.4536C47.9066 20.3728 48.7635 20.9946 49.73 20.9946H54.1792C55.3714 20.9948 55.8674 22.5213 54.9028 23.2222L51.3032 25.8364C50.5214 26.4044 50.1943 27.4115 50.4927 28.3306L51.8677 32.562C52.2361 33.6961 50.9378 34.6389 49.9731 33.938L46.3745 31.3237C45.5926 30.7556 44.5334 30.7556 43.7515 31.3237L40.1528 33.938C39.1882 34.6389 37.8898 33.6961 38.2583 32.562L39.6333 28.3306C39.9317 27.4115 39.6045 26.4044 38.8228 25.8364L35.2231 23.2222C34.2585 22.5213 34.7545 20.9947 35.9468 20.9946H40.396C41.3623 20.9945 42.2184 20.3726 42.5171 19.4536L43.8921 15.2222Z" stroke="#C5A576"/>
</g>
<path fill-rule="evenodd" clip-rule="evenodd" d="M33.3795 37.1831C33.6965 37.1831 33.9579 37.4419 33.9108 37.7554C33.7891 38.5646 33.4115 39.3201 32.8258 39.9058C32.1039 40.6274 31.1249 41.0326 30.1041 41.0327C29.0832 41.0327 28.1034 40.6277 27.3815 39.9058C26.7958 39.3201 26.4182 38.5646 26.2965 37.7554C26.2495 37.442 26.5109 37.1833 26.8278 37.1831C27.1446 37.1831 27.3956 37.4427 27.4625 37.7524C27.5712 38.2563 27.823 38.7231 28.193 39.0933C28.6997 39.6 29.3875 39.8853 30.1041 39.8853C30.8206 39.8852 31.5076 39.5999 32.0143 39.0933C32.3845 38.723 32.6361 38.2564 32.7448 37.7524C32.8116 37.4428 33.0627 37.1832 33.3795 37.1831Z" fill="#231717"/>
<mask id="path-7-inside-1_5551_16237" fill="white">
<path d="M24.5034 37.0017C24.5034 35.8971 25.3988 35.0017 26.5034 35.0017H33.5034C34.608 35.0017 35.5034 35.8971 35.5034 37.0017V42.5017C35.5034 44.987 33.4887 47.0017 31.0034 47.0017H29.0034C26.5181 47.0017 24.5034 44.987 24.5034 42.5017V37.0017Z"/>
</mask>
<path d="M24.5034 37.0017C24.5034 35.8971 25.3988 35.0017 26.5034 35.0017H33.5034C34.608 35.0017 35.5034 35.8971 35.5034 37.0017V42.5017C35.5034 44.987 33.4887 47.0017 31.0034 47.0017H29.0034C26.5181 47.0017 24.5034 44.987 24.5034 42.5017V37.0017Z" fill="black" stroke="black" stroke-width="11" mask="url(#path-7-inside-1_5551_16237)"/>
<ellipse cx="30.0034" cy="44.5017" rx="3.5" ry="1.5" fill="#E2759F"/>
</g>
<defs>
<filter id="filter0_d_5551_16237" x="3.08867" y="13.1207" width="23.2025" height="22.3117" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
<feFlood flood-opacity="0" result="BackgroundImageFix"/>
<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
<feOffset/>
<feGaussianBlur stdDeviation="0.375587"/>
<feComposite in2="hardAlpha" operator="out"/>
<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_5551_16237"/>
<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_5551_16237" result="shape"/>
</filter>
<filter id="filter1_d_5551_16237" x="33.4617" y="13.1207" width="23.2025" height="22.3117" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
<feFlood flood-opacity="0" result="BackgroundImageFix"/>
<feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
<feOffset/>
<feGaussianBlur stdDeviation="0.375587"/>
<feComposite in2="hardAlpha" operator="out"/>
<feColorMatrix type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0"/>
<feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_5551_16237"/>
<feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_5551_16237" result="shape"/>
</filter>
<clipPath id="clip0_5551_16237">
<rect width="60" height="60" rx="30" fill="white"/>
</clipPath>
</defs>
</svg>`;

const SUPER_HAPPY_SVG = `<svg width="60" height="60" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="59.9962" height="59.9962" rx="29.9981" fill="BGCOLOR"/>
<path d="M31.7417 25.0217H44.5913C44.88 25.0217 45.1145 25.2555 45.1147 25.5442V28.9651C45.1147 32.8023 42.0037 35.9133 38.1665 35.9133C34.3293 35.9133 31.2183 32.8023 31.2183 28.9651V25.5442C31.2185 25.2555 31.453 25.0217 31.7417 25.0217ZM27.8149 17.0881C28.5538 17.0882 29.1528 17.6871 29.1528 18.426C29.1528 19.1649 28.5538 19.7639 27.8149 19.7639C27.076 19.7639 26.4771 19.1649 26.4771 18.426C26.4771 17.6871 27.076 17.0881 27.8149 17.0881ZM48.5181 17.0881C49.2569 17.0883 49.856 17.6872 49.856 18.426C49.8559 19.1648 49.2568 19.7638 48.5181 19.7639C47.7792 19.7639 47.1802 19.1649 47.1802 18.426C47.1802 17.6871 47.7791 17.0881 48.5181 17.0881Z" fill="black" stroke="black" stroke-width="1.12669"/>
</svg>`;

// Order matches Figma layout (initials slot is handled separately)
export const UNIQUE_AVATAR_FACES: { key: string; svg: string }[] = [
  { key: "Alien", svg: ALIEN_SVG },
  { key: "Star Face", svg: STAR_FACE_SVG },
  { key: "Silly", svg: SILLY_SVG },
  { key: "Glasses Face", svg: GLASSES_FACE_SVG },
  { key: "Cute Face", svg: CUTE_FACE_SVG },
  { key: "Gap Tooth", svg: GAP_TOOTH_SVG },
  { key: "Super Happy", svg: SUPER_HAPPY_SVG },
  { key: "Sleepy", svg: SLEEPY_SVG },
];

export function getSvgWithColor(svgTemplate: string, color: string): string {
  return svgTemplate.replace("BGCOLOR", color);
}

export function getAvatarSvg(faceKey: string): string | null {
  const face = UNIQUE_AVATAR_FACES.find((f) => f.key === faceKey);
  return face ? face.svg : null;
}

type AvatarPickerProps = {
  selectedFace: string | null;
  selectedBackground: string;
  onFaceChange: (face: string | null) => void;
  onBackgroundChange: (color: string) => void;
  childInitials?: string;
};

type Tab = "Colors" | "Avatar";

// Renders a selection ring without touching overflow:hidden.
// The ring is on the outer wrapper; clipping is on the inner wrapper.
function RingWrapper({
  selected,
  onPress,
  children,
}: {
  selected: boolean;
  onPress: () => void;
  children: React.ReactNode;
}) {
  return (
    <TouchableOpacity
      onPress={onPress}
      activeOpacity={0.75}
      style={{
        padding: RING_GAP,
        borderRadius: CIRCLE_SIZE / 2 + RING_GAP + RING_WIDTH,
        borderWidth: RING_WIDTH,
        borderColor: selected ? "#6B7FC8" : "transparent",
      }}
    >
      {children}
    </TouchableOpacity>
  );
}

export function AvatarPicker({
  selectedFace,
  selectedBackground,
  onFaceChange,
  onBackgroundChange,
  childInitials = "?",
}: AvatarPickerProps) {
  const [activeTab, setActiveTab] = useState<Tab>("Colors");
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];

  const selectedBg = selectedBackground || DEFAULT_AVATAR_COLOR;

  // Split 9 items into rows of 3 for a guaranteed 3-column grid
  function toRows<T>(items: T[], perRow = 3): T[][] {
    const rows: T[][] = [];
    for (let i = 0; i < items.length; i += perRow) {
      rows.push(items.slice(i, i + perRow));
    }
    return rows;
  }

  const colorRows = toRows(AVATAR_COLORS);
  // Avatar grid: initials (null) + 8 faces = 9 items
  const avatarItems: (string | null)[] = [
    null,
    ...UNIQUE_AVATAR_FACES.map((f) => f.key),
  ];
  const avatarRows = toRows(avatarItems);

  return (
    <View
      style={{
        borderRadius: 16,
        borderWidth: 1,
        borderColor: theme.borderColor,
        marginBottom: 20,
        overflow: "hidden",
      }}
    >
      {/* Tab switcher */}
      <View
        style={{
          flexDirection: "row",
          padding: 6,
          gap: 6,
          backgroundColor: colorScheme === "dark" ? "#1c1c1e" : "#F3F4F6",
        }}
      >
        {(["Colors", "Avatar"] as Tab[]).map((tab) => (
          <TouchableOpacity
            key={tab}
            onPress={() => setActiveTab(tab)}
            style={{
              flex: 1,
              paddingVertical: 8,
              borderRadius: 10,
              alignItems: "center",
              backgroundColor: activeTab === tab ? "#8AA0CC" : "transparent",
            }}
          >
            <ThemedText
              style={{
                fontSize: 14,
                fontFamily: "Nunito-SemiBold",
                color: activeTab === tab ? "#fff" : theme.text,
              }}
            >
              {tab}
            </ThemedText>
          </TouchableOpacity>
        ))}
      </View>

      {/* Grid */}
      <View
        style={{
          paddingHorizontal: 16,
          paddingTop: 12,
          paddingBottom: 16,
          backgroundColor: colorScheme === "dark" ? "#1c1c1e" : "#fff",
        }}
      >
        {activeTab === "Colors"
          ? colorRows.map((row, ri) => (
              <View
                key={ri}
                style={{
                  flexDirection: "row",
                  justifyContent: "space-between",
                  marginBottom: ri < colorRows.length - 1 ? 14 : 0,
                }}
              >
                {row.map((color) => (
                  <RingWrapper
                    key={color}
                    selected={selectedBg === color}
                    onPress={() => onBackgroundChange(color)}
                  >
                    <View
                      style={{
                        width: CIRCLE_SIZE,
                        height: CIRCLE_SIZE,
                        borderRadius: CIRCLE_SIZE / 2,
                        backgroundColor: color,
                      }}
                    />
                  </RingWrapper>
                ))}
              </View>
            ))
          : avatarRows.map((row, ri) => (
              <View
                key={ri}
                style={{
                  flexDirection: "row",
                  justifyContent: "space-between",
                  marginBottom: ri < avatarRows.length - 1 ? 14 : 0,
                }}
              >
                {row.map((faceKey) => {
                  const isSelected = selectedFace === faceKey;
                  if (faceKey === null) {
                    return (
                      <RingWrapper
                        key="initials"
                        selected={isSelected}
                        onPress={() => onFaceChange(null)}
                      >
                        <View
                          style={{
                            width: CIRCLE_SIZE,
                            height: CIRCLE_SIZE,
                            borderRadius: CIRCLE_SIZE / 2,
                            backgroundColor: selectedBg,
                            alignItems: "center",
                            justifyContent: "center",
                          }}
                        >
                          <ThemedText
                            style={{
                              fontSize: 18,
                              fontFamily: "Nunito-Bold",
                              color: "#fff",
                            }}
                          >
                            {childInitials}
                          </ThemedText>
                        </View>
                      </RingWrapper>
                    );
                  }
                  const face = UNIQUE_AVATAR_FACES.find(
                    (f) => f.key === faceKey,
                  )!;
                  return (
                    <RingWrapper
                      key={faceKey}
                      selected={isSelected}
                      onPress={() => onFaceChange(faceKey)}
                    >
                      {/* overflow:hidden is on this inner View, completely
                          separate from the selection ring on RingWrapper */}
                      <View
                        style={{
                          width: CIRCLE_SIZE,
                          height: CIRCLE_SIZE,
                          borderRadius: CIRCLE_SIZE / 2,
                          overflow: "hidden",
                        }}
                      >
                        <SvgXml
                          xml={getSvgWithColor(face.svg, selectedBg)}
                          width={CIRCLE_SIZE}
                          height={CIRCLE_SIZE}
                        />
                      </View>
                    </RingWrapper>
                  );
                })}
              </View>
            ))}
      </View>
    </View>
  );
}
