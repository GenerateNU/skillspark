import React, { useRef, useState } from "react";
import {
	View,
	TextInput,
	TouchableOpacity,
	Alert,
	ScrollView,
	KeyboardAvoidingView,
	Platform,
	Keyboard,
	Pressable,
	Modal,
} from "react-native";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { AuthBackground } from "@/components/AuthBackground";
import { AppColors, Colors, FontSizes } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useQueryClient } from "@tanstack/react-query";
import {
	useCreateChild,
	useUpdateChild,
	useDeleteChild,
	getGetChildrenByGuardianIdQueryKey,
	useGetAllSchools,
	type School,
} from "@skillspark/api-client";
import { MONTHS } from "@/components/ChildProfileForm";
import { DEFAULT_AVATAR_COLOR } from "@/components/AvatarPicker";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorScreen } from "@/components/ErrorScreen";
import { Button } from "@/components/Button";

// Display labels shown in the UI (matching Figma)
const INTEREST_ROWS = [
	["Sports/Physical", "Music & Performance", "Languages"],
	["Personal Life Skills", "Academics", "Tech & Innovation"],
	["Small Group Tutoring", "Art/Creative Expression"],
];

// Maps display labels to valid Postgres category enum values
const DISPLAY_TO_VALUE: Record<string, string> = {
	"Sports/Physical": "sports",
	"Music & Performance": "music",
	"Languages": "language",
	"Personal Life Skills": "other",
	"Academics": "math",
	"Tech & Innovation": "technology",
	"Small Group Tutoring": "science",
	"Art/Creative Expression": "art",
};

// Maps enum values back to display labels
const VALUE_TO_DISPLAY: Record<string, string> = {
	sports: "Sports/Physical",
	music: "Music & Performance",
	language: "Languages",
	other: "Personal Life Skills",
	math: "Academics",
	technology: "Tech & Innovation",
	science: "Small Group Tutoring",
	art: "Art/Creative Expression",
};

const YEARS = Array.from({ length: 30 }, (_, i) =>
	String(new Date().getFullYear() - i),
);

type DropdownLayout = { x: number; y: number; width: number };

function DropdownModal({
	visible,
	onClose,
	layout,
	options,
	onSelect,
}: {
	visible: boolean;
	onClose: () => void;
	layout: DropdownLayout | null;
	options: string[];
	onSelect: (v: string) => void;
}) {
	if (!layout) return null;
	return (
		<Modal visible={visible} transparent animationType="none" onRequestClose={onClose}>
			<TouchableOpacity className="flex-1" activeOpacity={1} onPress={onClose}>
				<View
					className="absolute bg-white rounded-[10px] border border-[#E5E7EB] max-h-[220px] overflow-hidden"
					style={{
						top: layout.y,
						left: layout.x,
						width: layout.width,
						shadowColor: "#000",
						shadowOpacity: 0.12,
						shadowRadius: 8,
						shadowOffset: { width: 0, height: 2 },
						elevation: 8,
					}}
				>
					<ScrollView bounces={false}>
						{options.map((opt) => (
							<TouchableOpacity
								key={opt}
								className="px-4 py-3 border-b border-b-[#E5E7EB]"
								onPress={() => { onSelect(opt); onClose(); }}
							>
								<ThemedText>{opt}</ThemedText>
							</TouchableOpacity>
						))}
					</ScrollView>
				</View>
			</TouchableOpacity>
		</Modal>
	);
}

export default function AddChildScreen() {
	const router = useRouter();
	const params = useLocalSearchParams();
	const insets = useSafeAreaInsets();
	const theme = Colors.light;
	const { guardianId } = useAuthContext();
	const { t: translate } = useTranslation();
	const isEditing = !!params.id;

	const [firstName, setFirstName] = useState(
		params.name ? (params.name as string).split(" ")[0] : "",
	);
	const [lastName, setLastName] = useState(
		params.name ? (params.name as string).split(" ").slice(1).join(" ") : "",
	);
	const initialMonthStr = params.birth_month
		? MONTHS[parseInt(params.birth_month as string) - 1]
		: "";
	const [birthMonth, setBirthMonth] = useState(initialMonthStr);
	const [birthYear, setBirthYear] = useState((params.birth_year as string) || "");
	const [schoolId, setSchoolId] = useState((params.school_id as string) || "");

	// Stored values are backend enum strings — convert to display labels for the UI
	const rawInterests = Array.isArray(params.interests)
		? params.interests
		: params.interests
			? (params.interests as string).split(",").map((s) => s.trim()).filter(Boolean)
			: [];
	const initialInterests = rawInterests.map((v) => VALUE_TO_DISPLAY[v] ?? v);
	const [interests, setInterests] = useState<string[]>(initialInterests);
	const [isSubmitting, setIsSubmitting] = useState(false);

	const [showMonthDrop, setShowMonthDrop] = useState(false);
	const [showYearDrop, setShowYearDrop] = useState(false);
	const [showSchoolDrop, setShowSchoolDrop] = useState(false);
	const [monthDropLayout, setMonthDropLayout] = useState<DropdownLayout | null>(null);
	const [yearDropLayout, setYearDropLayout] = useState<DropdownLayout | null>(null);
	const [schoolDropLayout, setSchoolDropLayout] = useState<DropdownLayout | null>(null);
	const monthTriggerRef = useRef<View>(null);
	const yearTriggerRef = useRef<View>(null);
	const schoolTriggerRef = useRef<View>(null);

	const { data: schoolsData, isLoading: schoolsLoading } = useGetAllSchools();
	const schools: School[] = Array.isArray(schoolsData?.data) ? schoolsData.data : [];
	const selectedSchool = schools.find((s) => s.id === schoolId);

	const queryClient = useQueryClient();
	const createChildMutation = useCreateChild();
	const updateChildMutation = useUpdateChild();
	const deleteChildMutation = useDeleteChild();

	if (!guardianId) {
		return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
	}

	const openMonthDrop = () => {
		monthTriggerRef.current?.measure((_fx, _fy, width, height, px, py) => {
			setMonthDropLayout({ x: px, y: py + height + 4, width });
			setShowMonthDrop(true);
			setShowYearDrop(false);
			setShowSchoolDrop(false);
		});
	};

	const openYearDrop = () => {
		yearTriggerRef.current?.measure((_fx, _fy, width, height, px, py) => {
			setYearDropLayout({ x: px, y: py + height + 4, width });
			setShowYearDrop(true);
			setShowMonthDrop(false);
			setShowSchoolDrop(false);
		});
	};

	const openSchoolDrop = () => {
		schoolTriggerRef.current?.measure((_fx, _fy, width, height, px, py) => {
			setSchoolDropLayout({ x: px, y: py + height + 4, width });
			setShowSchoolDrop(true);
			setShowMonthDrop(false);
			setShowYearDrop(false);
		});
	};

	const toggleInterest = (item: string) =>
		setInterests((prev) =>
			prev.includes(item) ? prev.filter((i) => i !== item) : [...prev, item],
		);

	const handleSave = async () => {
		if (!firstName || !birthYear || !birthMonth || !schoolId) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}
		const name = [firstName, lastName].filter(Boolean).join(" ");
		setIsSubmitting(true);
		try {
			const childData = {
				name,
				birth_year: parseInt(birthYear, 10),
				birth_month: MONTHS.indexOf(birthMonth) + 1,
				guardian_id: guardianId,
				school_id: schoolId,
				interests: interests.map((label) => DISPLAY_TO_VALUE[label] ?? label),
				avatar_face: (params.avatar_face as string) || undefined,
				avatar_background: (params.avatar_background as string) || DEFAULT_AVATAR_COLOR,
			};
			let childId: string;
			if (isEditing) {
				await updateChildMutation.mutateAsync({ id: params.id as string, data: childData });
				childId = params.id as string;
			} else {
				const result = await createChildMutation.mutateAsync({ data: childData });
				childId = (result.data as any).id as string;
			}
			await queryClient.invalidateQueries({
				queryKey: getGetChildrenByGuardianIdQueryKey(guardianId),
			});
			router.push({
				pathname: "./edit-pic",
				params: {
					name,
					childId,
					avatar_face: (params.avatar_face as string) || "",
					avatar_background: (params.avatar_background as string) || DEFAULT_AVATAR_COLOR,
				},
			});
		} catch (error) {
			Alert.alert(
				translate("common.errorOccurred"),
				translate("childProfile.saveError") + " " + JSON.stringify(error),
			);
		} finally {
			setIsSubmitting(false);
		}
	};

	const handleDelete = () => {
		Alert.alert(
			translate("childProfile.deleteProfile"),
			translate("childProfile.deleteConfirm"),
			[
				{ text: translate("common.cancel"), style: "cancel" },
				{
					text: translate("payment.delete"),
					style: "destructive",
					onPress: async () => {
						setIsSubmitting(true);
						try {
							await deleteChildMutation.mutateAsync({ id: params.id as string });
							await queryClient.invalidateQueries({
								queryKey: getGetChildrenByGuardianIdQueryKey(guardianId),
							});
							router.back();
						} catch {
							Alert.alert(translate("common.errorOccurred"), translate("childProfile.deleteError"));
							setIsSubmitting(false);
						}
					},
				},
			],
		);
	};

	return (
		<View className="flex-1" style={{ paddingTop: insets.top }}>
			<AuthBackground />
			<Stack.Screen options={{ headerShown: false }} />
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				className="flex-1"
				keyboardVerticalOffset={0}
			>
				<ScrollView
					contentContainerStyle={{
						paddingHorizontal: 24,
						paddingBottom: insets.bottom + 32,
						paddingTop: 10,
					}}
					showsVerticalScrollIndicator={false}
					keyboardShouldPersistTaps="handled"
				>
					<Pressable onPress={Keyboard.dismiss}>
						{/* Header */}
						<View className="flex-row items-center justify-between mb-6">
							<TouchableOpacity
								onPress={() => router.back()}
								className="flex-row items-center gap-1"
								hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
							>
								<IconSymbol name="chevron.left" size={18} color={theme.text} />
								<ThemedText className="text-base font-nunito">
									{translate("onboarding.back")}
								</ThemedText>
							</TouchableOpacity>
							{isEditing && (
								<TouchableOpacity onPress={handleDelete}>
									<ThemedText className="font-nunito-semibold text-[#EF4444]">
										{translate("payment.delete")}
									</ThemedText>
								</TouchableOpacity>
							)}
						</View>

						{/* Title */}
						<ThemedText
							className="font-nunito-bold text-[#111] text-center mb-8"
							style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}
						>
							{translate("onboarding.setUpChild")}
						</ThemedText>

						{/* First Name */}
						<ThemedText className="font-nunito-semibold text-base text-[#111] mb-1">
							{translate("childProfile.firstName")}
						</ThemedText>
						<TextInput
							className="border border-black rounded-2xl px-4 h-[54px] bg-white text-base font-nunito text-[#11181C] mb-4"
							value={firstName}
							onChangeText={setFirstName}
							placeholderTextColor={AppColors.placeholderText}
						/>

						{/* Last Name */}
						<ThemedText className="font-nunito-semibold text-base text-[#111] mb-1">
							{translate("childProfile.lastName")}
						</ThemedText>
						<TextInput
							className="border border-black rounded-2xl px-4 h-[54px] bg-white text-base font-nunito text-[#11181C] mb-4"
							value={lastName}
							onChangeText={setLastName}
							placeholderTextColor={AppColors.placeholderText}
						/>

						{/* Birth Month + Year */}
						<View className="flex-row gap-3 mb-4">
							<View className="flex-1">
								<ThemedText className="font-nunito-semibold text-base text-[#111] mb-1">
									{translate("childProfile.birthMonth", { defaultValue: "Birth Month" })}
								</ThemedText>
								<View ref={monthTriggerRef}>
									<TouchableOpacity
										className="border border-black rounded-2xl px-4 h-[54px] bg-white flex-row items-center"
										onPress={openMonthDrop}
									>
										<ThemedText className="text-base font-nunito" style={{ color: theme.text }}>
											{birthMonth}
										</ThemedText>
									</TouchableOpacity>
								</View>
							</View>
							<View className="flex-1">
								<ThemedText className="font-nunito-semibold text-base text-[#111] mb-1">
									{translate("childProfile.birthYear", { defaultValue: "Birth Year" })}
								</ThemedText>
								<View ref={yearTriggerRef}>
									<TouchableOpacity
										className="border border-black rounded-2xl px-4 h-[54px] bg-white flex-row items-center"
										onPress={openYearDrop}
									>
										<ThemedText className="text-base font-nunito" style={{ color: theme.text }}>
											{birthYear}
										</ThemedText>
									</TouchableOpacity>
								</View>
							</View>
						</View>

						{/* School */}
						<ThemedText className="font-nunito-semibold text-base text-[#111] mb-1">
							{translate("childProfile.selectSchool", { defaultValue: "School" })}
						</ThemedText>
						<View ref={schoolTriggerRef} className="mb-8">
							<TouchableOpacity
								className="border border-black rounded-2xl px-4 h-[54px] bg-white flex-row items-center justify-between"
								onPress={openSchoolDrop}
								disabled={schoolsLoading}
							>
								<ThemedText
									className="text-base font-nunito"
									style={{ color: selectedSchool ? theme.text : AppColors.placeholderText }}
								>
									{schoolsLoading ? translate("childProfile.loadingSchools") : selectedSchool?.name ?? ""}
								</ThemedText>
								<IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
							</TouchableOpacity>
						</View>

						{/* Interests */}
						<ThemedText className="font-nunito text-base text-center text-[#111] leading-6 mb-5">
							{translate("childProfile.addInterestsHint", {
								defaultValue: "Add some interests. You can always change these later.",
							})}
						</ThemedText>
						<View className="gap-2 mb-8">
							{INTEREST_ROWS.map((row, rowIndex) => (
								<View key={rowIndex} className="flex-row justify-center gap-2">
									{row.map((item) => {
										const selected = interests.includes(item);
										return (
											<TouchableOpacity
												key={item}
												onPress={() => toggleInterest(item)}
												className="border rounded-full px-[14px] py-[5px]"
												style={{
													borderColor: selected ? "#7BAFD4" : "#000",
													backgroundColor: selected ? "#D9E4F5" : "#fff",
												}}
											>
												<ThemedText className="text-sm font-nunito text-[#11181C]">
													{item}
												</ThemedText>
											</TouchableOpacity>
										);
									})}
								</View>
							))}
						</View>

						<View className="items-center">
							<Button
								label={
									isSubmitting
										? translate("childProfile.saving")
										: translate("onboarding.submit", { defaultValue: "Submit" })
								}
								onPress={handleSave}
								disabled={isSubmitting}
							/>
						</View>
					</Pressable>
				</ScrollView>
			</KeyboardAvoidingView>

			<DropdownModal
				visible={showMonthDrop}
				onClose={() => setShowMonthDrop(false)}
				layout={monthDropLayout}
				options={MONTHS}
				onSelect={setBirthMonth}
			/>
			<DropdownModal
				visible={showYearDrop}
				onClose={() => setShowYearDrop(false)}
				layout={yearDropLayout}
				options={YEARS}
				onSelect={setBirthYear}
			/>
			<DropdownModal
				visible={showSchoolDrop}
				onClose={() => setShowSchoolDrop(false)}
				layout={schoolDropLayout}
				options={schools.map((s) => s.name)}
				onSelect={(name) => {
					const school = schools.find((s) => s.name === name);
					if (school) setSchoolId(school.id);
				}}
			/>
		</View>
	);
}
