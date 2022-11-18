import {AfterViewInit, Component, Input, OnInit, ViewChild} from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {TableDataService} from '../../services/data.service';
import {SelectionModel} from '@angular/cdk/collections';
import {MatSort} from '@angular/material/sort';

@Component({
    selector: 'app-table',
    templateUrl: './table.component.html',
    styleUrls: ['./table.component.scss'],
})

export class TableComponent implements OnInit, AfterViewInit {
    @Input() displayedColumns: any[];
    @Input() isHidden: boolean;
    @Input() selection: SelectionModel<any>

    @ViewChild(MatSort) sort: MatSort;

    dataSources: MatTableDataSource<any>;

    constructor(public tableDataService: TableDataService) { }

    ngAfterViewInit() {
        this.dataSources.sort = this.sort;
    }

    ngOnInit() {
        this.dataSources = new MatTableDataSource(this.tableDataService.changeMapToArray(this.tableDataService.retrieveItemsFromCache()));
    }

    applyFilter(event: Event) {
        const filterValue = (event.target as HTMLInputElement).value;
        this.dataSources.filter = filterValue.trim().toLowerCase();
    }

    allRowsSelected() {
        const totalRowSelected = this.selection.selected.length;
        const totalRow = this.dataSources.data.length;
        return totalRowSelected === totalRow;
    }

    masterToggle() {
        this.allRowsSelected() ?
            this.selection.clear() :
            this.dataSources.data.forEach(element => this.selection.select(element));
    }
}
