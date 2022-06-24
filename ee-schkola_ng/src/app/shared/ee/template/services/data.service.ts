import {Injectable} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {FormService} from './form.service';

@Injectable({ providedIn: 'root' })
export class TableDataService {
    itemName = '';
    items: Array<any> = [];

    constructor(private templateService: FormService, private _router: Router, private _route: ActivatedRoute) {
    }

    addItemToTableArray(items: Array<any>) {
        this.retrieveItemFromCache();
        this.items.push(items);
        localStorage.setItem(this.itemName, JSON.stringify(this.items));
    }

    retrieveItemFromCache() {
        this.items = JSON.parse(localStorage.getItem(this.itemName));
        return this.items;
    }

    clearItems() {
        localStorage.setItem(this.itemName, JSON.stringify([]));
        this.items = [];
        window.location.reload();
    }

    removeItem(index) {
        this.retrieveItemFromCache();
        this.items.splice(index, 1);
        localStorage.setItem(this.itemName, JSON.stringify(this.items));
        window.location.reload();
    }

    editItems(index) {
        this._router.navigate(['edit' , index], {relativeTo: this._route});
    }
}
