import { getLocationById, getManagerByOrgId, type Manager, type Organization, type Location } from "@skillspark/api-client";


export function loadData(orgFromState : Organization, setManager : (man : Manager) => void, setOrgLocation : (loc : Location) => void, setLoadingMgr : (bool : boolean) => void) {
    async function fetchManager(): Promise<void> {
      try {
        const res = await getManagerByOrgId(orgFromState.id);
        if (res.status === 200) setManager(res.data as Manager);
      } catch (e) {
        console.error(e);
      } finally {
        setLoadingMgr(false);
      }
    }

    async function fetchLocation(): Promise<void> {
      if (!orgFromState.location_id) return;
      try {
        const res = await getLocationById(orgFromState.location_id as string);
        if (res.status === 200 || res.status === 201) setOrgLocation(res.data as Location);
      } catch (e) {
        console.error(e);
      }
    }

    fetchManager();
    fetchLocation();
}
