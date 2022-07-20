import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FormService} from './form.service';

@Injectable({ providedIn: 'root' })
export class TableDataService {

    public isEdit: boolean;
    public itemIndex: number;
    itemName = '';
    items: Map<any, any> = new Map();

    constructor(private templateService: FormService, private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Object, id: string) {
        this.items = this.retrieveItemFromCache();
        this.items.set(id, items);
        localStorage.setItem(this.itemName, JSON.stringify(Array.from(this.items.entries())));
        localStorage.myMap = JSON.stringify(Array.from(this.items.entries()));
    }

    retrieveItemFromCache() {
        this.items = new Map(JSON.parse(localStorage.myMap));
        return this.items;
    }

    clearItems() {
        this.items.clear();
        localStorage.setItem(this.itemName, JSON.stringify(this.items));
        window.location.reload();
    }

    removeItem(id) {
        this.retrieveItemFromCache();
        this.items.delete(id);
        localStorage.setItem(this.itemName, JSON.stringify(this.items));
        window.location.reload();
    }

    editItems(index) {
        this._router.navigate(['edit' , index], {relativeTo: this._route});
    }

    loadEnumElement(enumElement: any) {
        let tempArray = [];
        Object.keys(enumElement).map((element, index) => {
            tempArray.push(enumElement[index]);
        })
        tempArray = tempArray.filter((item) => item);
        return tempArray;
    }
}
