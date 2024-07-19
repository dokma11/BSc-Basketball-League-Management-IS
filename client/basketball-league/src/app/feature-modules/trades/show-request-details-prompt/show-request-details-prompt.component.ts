import { trigger, transition, style, animate, state } from '@angular/animations';
import { Component, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MatDialog } from '@angular/material/dialog';
import { MatSelect } from '@angular/material/select';
import { MatSnackBar } from '@angular/material/snack-bar';
import { faPlus, faMinus } from '@fortawesome/free-solid-svg-icons';
import { Observable, of, Subject, startWith, map, take, takeUntil } from 'rxjs';
import { AssetChoosingFormComponent } from '../asset-choosing-form/asset-choosing-form.component';
import { ProposeTradeFormComponent } from '../propose-trade-form/propose-trade-form.component';

@Component({
  selector: 'app-show-request-details-prompt',
  templateUrl: './show-request-details-prompt.component.html',
  styleUrls: ['./show-request-details-prompt.component.css'],
  animations: [
      trigger("fadeIn", [
        transition(":enter", [
            style({ opacity: 0, transform: "translateX(-40px)" }),
            animate(
                "0.5s ease",
                style({ opacity: 1, transform: "translateX(0)" }),
            ),
        ]),
      ]),
      trigger('buttonState', [
        state('clicked', style({
          transform: 'scale(0.9)',
          opacity: 0.5
        })),
        transition('* => clicked', [
          animate('200ms')
        ]),
        transition('clicked => idle', [
          animate('200ms')
        ])
      ]),
  ],
})
export class ShowRequestDetailsPromptComponent {
  buttonState: string = 'idle';
  removeAssetButtonState: string = 'idle';
  addPartnersAssetButtonState: string = 'idle';
  addYoursAssetButtonState: string = 'idle';
  focused: string = '';
  private ownDialogRef: any;
  public teams: string[] = [
    'Brooklyn Nets',
    'Golden State Warriors',
    'Los Angeles Lakers',
    'Los Angeles Clippers',
    'New Orleans Pelicans',
    'New York Knicks',
    'Oklahoma City Thunder',
    'San Antonio Spurs',
    'Boston Celtics',
    'Denver Nuggets',
    'Minnesota Timberwolves',
    'Cleveland Cavaliers',
    'Philadelphia 76ers',
    'Phoenix Suns',
    'Sacramento Kings',
    'Indiana Pacers',
    'Dallas Mavericks',
    'Miami Heat',
    'Orlando Magic',
    'Chicago Bulls',
    'Atlanta Hawks',
    'Toronto Raptors',
    'Charlotte Hornets',
    'Washington Wizards',
    'Detroit Pistons',
    'Utah Jazz',
    'Houston Rockets',
    'Memphis Grizzlies',
    'Portland Trail Blazers',
    'Milwaukee Bucks'
  ];

  public teamCtrl: FormControl<string | null> = new FormControl<string | null>('');
  public teamFilterCtrl: FormControl<string | null> = new FormControl<string | null>('');

  public filteredTeams: Observable<string[]> = of(this.teams);

  @ViewChild('singleSelect', { static: true }) singleSelect: MatSelect | undefined;

  protected _onDestroy = new Subject<void>();

  constructor(private snackBar: MatSnackBar,
              private dialogRef: MatDialogRef<ProposeTradeFormComponent>,
              private dialogRefAsset: MatDialogRef<AssetChoosingFormComponent>,
              private dialog: MatDialog,) {
  }

  ngOnInit(): void {
    this.teamCtrl.setValue('');

    this.filteredTeams = this.teamFilterCtrl.valueChanges.pipe(
      startWith(''),
      map(value => this.filterTeams(value))
    );
  }

  private filterTeams(value: any): string[] {
    const filterValue = value?.toLowerCase() || '';
    return this.teams.filter(team => team.toLowerCase().includes(filterValue));
  }

  ngAfterViewInit() {
    this.setInitialValue();
  }

  ngOnDestroy() {
    this._onDestroy.next();
    this._onDestroy.complete();
  }

  protected setInitialValue() {
    this.filteredTeams
      .pipe(take(1), takeUntil(this._onDestroy))
      .subscribe(() => {
        this.singleSelect!.compareWith = (a: string, b: string) => a.toLowerCase() === b.toLowerCase();
      });
  }

  // Izmeni samo da bude trejd kao
  proposeTradeForm = new FormGroup({
    name: new FormControl('', [Validators.required]),
    description: new FormControl('', [Validators.required]),
    //duration: new FormControl('', [Validators.required]),
    occurrenceTime: new FormControl(null, [Validators.required]),
    occurrenceDate: new FormControl(null, [Validators.required]),
    //guide: new FormControl('', [Validators.required]),
    capacity: new FormControl('', [Validators.required]),
    picturePath: new FormControl('', [Validators.required]),
    category: new FormControl('', [Validators.required]),
  });

  closeButtonClicked() {
    this.buttonState = 'clicked';
    setTimeout(() => { this.buttonState = 'idle'; }, 200);
    this.dialogRef.close();
  }

  addPartnersAssetButtonClicked(): void {
    this.addPartnersAssetButtonState = 'clicked';
    setTimeout(() => { this.addPartnersAssetButtonState = 'idle'; }, 200);
    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        // TODO: Ovde treba proslediti tim da se zna cija imovina da se prikaze
    });

    if (this.dialogRefAsset) {
      this.dialogRefAsset.afterClosed().subscribe((result: any) => {
        // TODO: Ovde treba dodati osvezavanje liste kada se odabere nova imovina za trejdovanje
      });
    }
  }

  addYoursAssetButtonClicked(): void {
    this.addYoursAssetButtonState = 'clicked';
    setTimeout(() => { this.addYoursAssetButtonState = 'idle'; }, 200);
    this.dialogRefAsset = this.dialog.open(AssetChoosingFormComponent, {
        // TODO: Ovde treba proslediti tim da se zna cija imovina da se prikaze
    });

    if (this.dialogRefAsset) {
      this.dialogRefAsset.afterClosed().subscribe((result: any) => {
        // TODO: Ovde treba dodati osvezavanje liste kada se odabere nova imovina za trejdovanje
      });
    }
  }

  removeAssetButtonClicked(): void {
    // Ovde samo proveri da li je sve okej sto se tice same animacije
    this.removeAssetButtonState = 'clicked';
    setTimeout(() => { this.removeAssetButtonState = 'idle'; }, 200);
  }

  showNotification(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'right',
      verticalPosition: 'bottom',
    });
  }

  overviewClicked(){
    this.dialogRef.close();
  }

  faPlus = faPlus;
  faMinus = faMinus;
}
