import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AssetChoosingFormComponent } from './asset-choosing-form.component';

describe('AssetChoosingFormComponent', () => {
  let component: AssetChoosingFormComponent;
  let fixture: ComponentFixture<AssetChoosingFormComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AssetChoosingFormComponent]
    });
    fixture = TestBed.createComponent(AssetChoosingFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
