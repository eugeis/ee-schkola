import {Component, Input, OnInit} from '@angular/core';
import {TemplateService} from '../../services/template.service';
import {Router} from '@angular/router';

@Component({
    selector: 'app-form',
    templateUrl: './form.component.html',
    styleUrls: ['./form.component.scss'],
    providers: [TemplateService],
})

export class FormComponent implements OnInit {

    @Input() formElement: any[][];
    public isEdit: boolean;
    public itemIndex: number;

    constructor(public templateService: TemplateService, private _router: Router) { }

    ngOnInit(): void {
        this.templateService.initElement(this.formElement);
        const currentUrl = this._router.url;
        currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase() !== 'new' ? this.isEdit = true : this.isEdit = false;
        this.itemIndex = Number(currentUrl.substring(currentUrl.lastIndexOf('/') + 1).toLowerCase());
        this.templateService.loadElement(this.itemIndex);
    }
}
