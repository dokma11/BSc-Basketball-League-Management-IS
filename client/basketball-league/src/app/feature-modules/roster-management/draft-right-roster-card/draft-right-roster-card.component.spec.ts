import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DraftRightRosterCardComponent } from './draft-right-roster-card.component';

describe('DraftRightRosterCardComponent', () => {
  let component: DraftRightRosterCardComponent;
  let fixture: ComponentFixture<DraftRightRosterCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DraftRightRosterCardComponent]
    });
    fixture = TestBed.createComponent(DraftRightRosterCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
