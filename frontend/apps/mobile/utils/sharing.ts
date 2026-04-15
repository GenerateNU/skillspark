import { Share } from "react-native";
import * as Linking from "expo-linking";

export async function shareEventLink(id: string, title?: string) {
  const url = Linking.createURL(`event/${id}`);
  await Share.share({
    message: title ? `Check out "${title}" on SkillSpark: ${url}` : url,
    url,
  });
}

export async function shareOrgLink(id: string, name?: string) {
  const url = Linking.createURL(`org/${id}`);
  await Share.share({
    message: name ? `Check out ${name} on SkillSpark: ${url}` : url,
    url,
  });
}
