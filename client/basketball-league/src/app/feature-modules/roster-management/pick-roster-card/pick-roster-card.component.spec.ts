import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PickRosterCardComponent } from './pick-roster-card.component';

describe('PickRosterCardComponent', () => {
  let component: PickRosterCardComponent;
  let fixture: ComponentFixture<PickRosterCardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PickRosterCardComponent]
    });
    fixture = TestBed.createComponent(PickRosterCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
