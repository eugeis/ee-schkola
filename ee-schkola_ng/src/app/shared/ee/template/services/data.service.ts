import {Injectable} from '@angular/core';
import {Router} from '@angular/router';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public isSearch: boolean;

    filterValue: string;
    itemName = '';

    constructor(protected _router?: Router) {
    }

    loadEnumElement(enumElement: any, parentName?: string, elementName?: string) {
        const tempArray = [];
        Object.keys(enumElement).map((element, index) => {
            if (enumElement[index] !== undefined) {
                tempArray.push(parentName + '.' + elementName + '.' + enumElement[index]);
            }
        })
        return tempArray;
    }

    checkSearchRoute() {
        const currentUrl = this._router.url;
        currentUrl.includes('search') ? this.isSearch = true : this.isSearch = false;
    }
}

declare global {
    interface Window {
        tableDataService: TableDataService;
    }
}
window.tableDataService = new TableDataService();
