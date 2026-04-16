import { View } from "react-native";
import { SvgXml } from "react-native-svg";

const LOGO_SVG = `<svg width="147" height="200" viewBox="0 0 147 200" fill="none" xmlns="http://www.w3.org/2000/svg">
<g opacity="0.59">
<g opacity="0.27">
<path d="M89.0811 67.8166L48.3281 21.8405C38.573 10.8351 21.0265 12.1061 12.9594 24.4025C7.27644 33.0647 8.30026 44.4985 15.4317 52.0133L110.288 151.968C123.993 166.409 113.756 190.237 93.8466 190.237C87.8351 190.237 82.0698 187.849 77.819 183.598L28.9689 134.748" stroke="url(#paint0_linear_5475_5621)" stroke-width="18.6021" stroke-linecap="round"/>
<path d="M108.28 73.7524C109.03 71.7254 111.897 71.7254 112.647 73.7524L115.366 81.1001C115.602 81.7374 116.105 82.2398 116.742 82.4756L124.09 85.1945C126.117 85.9446 126.117 88.8115 124.09 89.5615L116.742 92.2804C116.105 92.5162 115.602 93.0187 115.366 93.6559L112.647 101.004C111.897 103.031 109.03 103.031 108.28 101.004L105.562 93.6559C105.326 93.0187 104.823 92.5162 104.186 92.2804L96.8383 89.5615C94.8113 88.8115 94.8113 85.9446 96.8383 85.1945L104.186 82.4756C104.823 82.2398 105.326 81.7374 105.562 81.1001L108.28 73.7524Z" fill="#99C0EE"/>
</g>
</g>
<defs>
<linearGradient id="paint0_linear_5475_5621" x1="53.0336" y1="86.213" x2="62.8474" y2="222.951" gradientUnits="userSpaceOnUse">
<stop offset="0.195258" stop-color="#99C0EE"/>
<stop offset="1" stop-color="#8F3F95"/>
</linearGradient>
</defs>
</svg>`;

const scale = 1.2;
const LOGO_WIDTH = 147 * scale;
const LOGO_HEIGHT = 200 * scale;
const RATIO_HIDDEN = 3 / 8;

type LogoBgWrapperProps = {
  children: React.ReactNode;
  className?: string;
  verticalOffset?: number;
};

/**
 * Wraps children with a decorative background featuring the SkillSpark logo
 * half-cropped on the left and right edges, aligned at the same vertical position.
 */
export default function LogoBgWrapper({
  children,
  className,
  verticalOffset = 0,
}: LogoBgWrapperProps) {
  return (
    <View className={`overflow-hidden ${className ?? ""}`}>
      {/* Left logo — half off the left edge, mirrored horizontally */}
      <View
        style={{
          position: "absolute",
          left: -LOGO_WIDTH * RATIO_HIDDEN,
          top: verticalOffset,
        }}
        pointerEvents="none"
      >
        <SvgXml xml={LOGO_SVG} width={LOGO_WIDTH} height={LOGO_HEIGHT} />
      </View>

      {/* Right logo — half off the right edge */}
      <View
        style={{
          position: "absolute",
          right: -LOGO_WIDTH * RATIO_HIDDEN,
          top: verticalOffset,
        }}
        pointerEvents="none"
      >
        <SvgXml xml={LOGO_SVG} width={LOGO_WIDTH} height={LOGO_HEIGHT} />
      </View>

      {children}
    </View>
  );
}
