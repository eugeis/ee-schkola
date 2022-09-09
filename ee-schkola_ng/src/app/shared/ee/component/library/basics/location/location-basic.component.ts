import {Location} from '../library/LibraryApiBase'

import {Component, Input, OnInit} from '@angular/core';

@Component({
  selector: 'app-location',
  templateUrl: './location.component.html',
  styleUrls: ['./location.component.scss'],
})

export class LocationComponent implements OnInit {

    @Input() location: Location = new Location();

    ngOnInit(): void {
        if (this.location === undefined) {
            this.location = {shelf: '', fold: ''};
        }
    }
}




