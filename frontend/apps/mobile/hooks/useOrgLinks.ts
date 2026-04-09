import { useCallback } from "react";
import * as WebBrowser from "expo-web-browser";
import type { OrgLink } from "@skillspark/api-client";

export function useOrgLinks(links: OrgLink[]) {
  const openLink = useCallback(async (href: string) => {
    await WebBrowser.openBrowserAsync(href);
  }, []);

  return { openLink, hasLinks: links.length > 0 };
}
